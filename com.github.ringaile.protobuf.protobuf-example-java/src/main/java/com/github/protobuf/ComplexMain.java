package com.github.protobuf;

import example.complex.Complex;

import java.util.Arrays;

public class ComplexMain {

    public static void main(String[] args) {
        System.out.println("Complex example");

        Complex.DummyMessage oneDummy = newDummyMessage( 12, "one");

        Complex.ComplexMessage.Builder builder = Complex.ComplexMessage.newBuilder();
        builder.setOneDummy(oneDummy);

        //repeated field
        builder.addMultipleDummy(newDummyMessage(45, "second"));
        builder.addAllMultipleDummy(Arrays.asList(newDummyMessage(45, "third"), newDummyMessage(45, "4th")));

        Complex.ComplexMessage message = builder.build();
        System.out.println(message);
    }

    public static Complex.DummyMessage newDummyMessage(int id, String name) {
        Complex.DummyMessage.Builder builder = Complex.DummyMessage.newBuilder();
        return builder.setName(name)
                .setId(id)
                .build();
    }
}
