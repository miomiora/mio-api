package com.mio.mioapigateway.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.51.1)",
    comments = "Source: mio-api.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class mioApiGrpc {

  private mioApiGrpc() {}

  public static final String SERVICE_NAME = "mioApi";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<Request,
      Response> getGetSecretKeyMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetSecretKey",
      requestType = Request.class,
      responseType = Response.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<Request,
      Response> getGetSecretKeyMethod() {
    io.grpc.MethodDescriptor<Request, Response> getGetSecretKeyMethod;
    if ((getGetSecretKeyMethod = mioApiGrpc.getGetSecretKeyMethod) == null) {
      synchronized (mioApiGrpc.class) {
        if ((getGetSecretKeyMethod = mioApiGrpc.getGetSecretKeyMethod) == null) {
          mioApiGrpc.getGetSecretKeyMethod = getGetSecretKeyMethod =
              io.grpc.MethodDescriptor.<Request, Response>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetSecretKey"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  Request.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  Response.getDefaultInstance()))
              .setSchemaDescriptor(new mioApiMethodDescriptorSupplier("GetSecretKey"))
              .build();
        }
      }
    }
    return getGetSecretKeyMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static mioApiStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<mioApiStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<mioApiStub>() {
        @Override
        public mioApiStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new mioApiStub(channel, callOptions);
        }
      };
    return mioApiStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static mioApiBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<mioApiBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<mioApiBlockingStub>() {
        @Override
        public mioApiBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new mioApiBlockingStub(channel, callOptions);
        }
      };
    return mioApiBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static mioApiFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<mioApiFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<mioApiFutureStub>() {
        @Override
        public mioApiFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new mioApiFutureStub(channel, callOptions);
        }
      };
    return mioApiFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class mioApiImplBase implements io.grpc.BindableService {

    /**
     */
    public void getSecretKey(Request request,
                             io.grpc.stub.StreamObserver<Response> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetSecretKeyMethod(), responseObserver);
    }

    @Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetSecretKeyMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                Request,
                Response>(
                  this, METHODID_GET_SECRET_KEY)))
          .build();
    }
  }

  /**
   */
  public static final class mioApiStub extends io.grpc.stub.AbstractAsyncStub<mioApiStub> {
    private mioApiStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @Override
    protected mioApiStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new mioApiStub(channel, callOptions);
    }

    /**
     */
    public void getSecretKey(Request request,
                             io.grpc.stub.StreamObserver<Response> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetSecretKeyMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class mioApiBlockingStub extends io.grpc.stub.AbstractBlockingStub<mioApiBlockingStub> {
    private mioApiBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @Override
    protected mioApiBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new mioApiBlockingStub(channel, callOptions);
    }

    /**
     */
    public Response getSecretKey(Request request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetSecretKeyMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class mioApiFutureStub extends io.grpc.stub.AbstractFutureStub<mioApiFutureStub> {
    private mioApiFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @Override
    protected mioApiFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new mioApiFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<Response> getSecretKey(
        Request request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetSecretKeyMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_SECRET_KEY = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final mioApiImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(mioApiImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @Override
    @SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_SECRET_KEY:
          serviceImpl.getSecretKey((Request) request,
              (io.grpc.stub.StreamObserver<Response>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @Override
    @SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class mioApiBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    mioApiBaseDescriptorSupplier() {}

    @Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return MioApiProto.getDescriptor();
    }

    @Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("mioApi");
    }
  }

  private static final class mioApiFileDescriptorSupplier
      extends mioApiBaseDescriptorSupplier {
    mioApiFileDescriptorSupplier() {}
  }

  private static final class mioApiMethodDescriptorSupplier
      extends mioApiBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    mioApiMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (mioApiGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new mioApiFileDescriptorSupplier())
              .addMethod(getGetSecretKeyMethod())
              .build();
        }
      }
    }
    return result;
  }
}
