package com.mio.mioapigateway.client;

import com.mio.mioapigateway.proto.Request;
import com.mio.mioapigateway.proto.Response;
import com.mio.mioapigateway.proto.mioApiGrpc;
import io.grpc.Channel;
import io.grpc.StatusRuntimeException;

public class ClientRPC {
    private final mioApiGrpc.mioApiBlockingStub blockingStub;

    public ClientRPC(Channel channel) {
        blockingStub = mioApiGrpc.newBlockingStub(channel);
    }

    public Response getSecretKey(String accessKey) {
        Request request = Request.newBuilder().setRequestData(accessKey).build();
        Response response;
        try {
            response = blockingStub.getSecretKey(request);
        } catch (StatusRuntimeException e) {
            System.out.println("RPC failed: {0}"+ e.getStatus());
            return null;
        }
        return response;
    }

//    public static void main(String[] args) throws Exception {
//        String user = "world";
//        String target = "localhost:2175";
//        ManagedChannel channel = Grpc.newChannelBuilder(target, InsecureChannelCredentials.create())
//                .build();
//        try {
//            ClientRPC client = new ClientRPC(channel);
//            client.getSecretKey(user);
//        } finally {
//            // ManagedChannels use resources like threads and TCP connections. To prevent leaking these
//            // resources the channel should be shut down when it will no longer be used. If it may be used
//            // again leave it running.
//            channel.shutdownNow().awaitTermination(5, TimeUnit.SECONDS);
//        }
//    }


}
