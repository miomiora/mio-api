package com.mio.mioapigateway;

import com.mio.mioapigateway.client.ClientRPC;
import com.mio.mioapigateway.proto.Response;
import com.mio.mioapigateway.utils.SignUtils;
import io.grpc.Grpc;
import io.grpc.InsecureChannelCredentials;
import io.grpc.ManagedChannel;
import lombok.extern.slf4j.Slf4j;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.core.Ordered;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.server.reactive.ServerHttpRequest;
import org.springframework.http.server.reactive.ServerHttpResponse;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.util.Arrays;
import java.util.List;

@Component
@Slf4j
public class CustomGlobalFilter implements GlobalFilter, Ordered {

    private static final List<String> IP_WHITE_LIST = Arrays.asList("127.0.0.1");
    ManagedChannel channel = Grpc.newChannelBuilder("127.0.0.1:2175", InsecureChannelCredentials.create())
            .build();
    ClientRPC client = new ClientRPC(channel);

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        ServerHttpRequest request = exchange.getRequest();
        ServerHttpResponse response = exchange.getResponse();
        // 请求日志
        log.info("请求标识："+ request.getId());
        log.info("请求路径："+ request.getPath().value());
        log.info("请求方法："+ request.getMethodValue());
        log.info("请求参数："+ request.getQueryParams());
        String sourceAddress = request.getLocalAddress().getHostString();
        log.info("请求地址："+sourceAddress);
        log.info("请求来源："+ request.getRemoteAddress());

        // 黑白名单
        if (!IP_WHITE_LIST.contains(sourceAddress)) {
            response.setStatusCode(HttpStatus.FORBIDDEN);
            return response.setComplete();
        }

        HttpHeaders headers = request.getHeaders();
        String accessKey = headers.getFirst("accessKey");
        String nonce = headers.getFirst("nonce");
        String timestamp = headers.getFirst("timestamp");
        String sign = headers.getFirst("sign");
        String body = headers.getFirst("body");
        if (accessKey==null || nonce == null || timestamp == null || sign == null || body == null) {
            return handleInvokeError(response);
        }


        // 用户鉴权
        Long currentTime = System.currentTimeMillis() / 1000;
        Long FIVE_MINUTES = 60 * 5L;
        // 是否超过5分钟
        if ((currentTime - Long.parseLong(timestamp)) >= FIVE_MINUTES) {
            return handleNoAuth(response);
        }
        // 从gRPC服务器验证accessKey是否有效
        Response secretKey = client.getSecretKey(accessKey);
        String genSignStr = SignUtils.genSign(body, secretKey.getResponseData());
        if (!sign.equals(genSignStr)) {
            return handleNoAuth(response);
        }

        // 请求转发
        Mono<Void> filter = chain.filter(exchange);
//        log.info();

        // todo 调用成功接口次数+1

        // 调用失败
        if (response.getStatusCode() == HttpStatus.OK) {

        } else {
            return handleInvokeError(response);
        }

//        log.info("================================================");
        return filter;
    }

    @Override
    public int getOrder() {
        return -1;
    }

    public Mono<Void> handleNoAuth(ServerHttpResponse response) {
        response.setStatusCode(HttpStatus.FORBIDDEN);
        return response.setComplete();
    }

    public Mono<Void> handleInvokeError(ServerHttpResponse response) {
        response.setStatusCode(HttpStatus.INTERNAL_SERVER_ERROR);
        return response.setComplete();
    }

}
