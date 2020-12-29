package com.github.protobuf;

import example.simple.Simple.SimpleMessage;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.Arrays;

public class SimpleMain {

    public static void main(String[] args) throws IOException {
        System.out.println("Hello world!");

        SimpleMessage.Builder builder = SimpleMessage.newBuilder();

        builder.setId(42)
                .setIsSimple(true)
                .setName("My simple message name");

        builder.addSampleList(1);
        builder.addSampleList(2);

        builder.addAllSampleList(Arrays.asList(3,4,5));

        System.out.println(builder.toString());

        SimpleMessage message = builder.build();

        FileOutputStream outputStream = new FileOutputStream("simple_message.bin");
        message.writeTo(outputStream);
        outputStream.close();

        //send as byte array
        byte[] bytes = message.toByteArray();

        FileInputStream fileInputStream = new FileInputStream("simple_message.bin");
        SimpleMessage messageFromFile = SimpleMessage.parseFrom(fileInputStream);
        System.out.println(messageFromFile);
    }
}
