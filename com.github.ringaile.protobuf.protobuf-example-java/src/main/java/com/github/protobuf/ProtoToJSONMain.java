package com.github.protobuf;

import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.util.JsonFormat;
import example.simple.Simple;

import java.util.Arrays;

public class ProtoToJSONMain {

    public static void main(String[] args) throws InvalidProtocolBufferException {
        System.out.println("Hello world!");

        Simple.SimpleMessage.Builder builder = Simple.SimpleMessage.newBuilder();

        builder.setId(42)
                .setIsSimple(true)
                .setName("My simple message name");

        builder.addSampleList(1);
        builder.addSampleList(2);

        builder.addAllSampleList(Arrays.asList(3,4,5));

        String jsonString = JsonFormat.printer().print(builder);
        System.out.println(jsonString);

        Simple.SimpleMessage.Builder builder2 = Simple.SimpleMessage.newBuilder();
        JsonFormat.parser().merge(jsonString, builder2);

        System.out.println(builder2.build());
    }
}
