/*
 * Copyright (c) 2017, CipherGateway and/or its affiliates. All rights  reserved.
 *
 */
package text;

import java.util.Scanner;

import static java.lang.System.in;

/**
 * @author kyle
 */
public class StringTest {

  public static void main(String[] args) {
    // Scanner scanner = new Scanner(in);
    //int n = scanner.nextInt();
    String a = "Programming";
    String b = new String("Programming");
    String c = "Program" + "ming";

    System.out.println(a==b);
    System.out.println(a==c);
    System.out.println(a.equals(b));
    System.out.println(a.equals(c));
    System.out.println(a.intern() == b.intern());


  }
}
