package com.youstars.sgsocket;

import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelFutureListener;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.socket.SocketChannel;
import io.netty.handler.codec.http.HttpObjectAggregator;
import io.netty.handler.codec.http.HttpServerCodec;
import io.netty.handler.stream.ChunkedWriteHandler;

public class SmartGateSocketServerHandler extends ChannelInitializer<SocketChannel> {

    @Override
    protected void initChannel(SocketChannel ch) throws Exception {
        ch.pipeline().addLast(new HttpServerCodec()).addLast(new HttpObjectAggregator(64 * 1024))
                .addLast(new ChunkedWriteHandler()).addLast(new HttpRequestHandler()).addLast(new WebSocketFrameHandler());
        ch.closeFuture().addListener(new ChannelFutureListener() {
            public void operationComplete(ChannelFuture future) throws Exception {
                ChannelHub.removeChannel(future.channel());
            }
        });
    }
}