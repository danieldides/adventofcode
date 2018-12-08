package week1;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class Application {

    public static void main(String[] args) throws IOException {
        System.out.println(Files.lines(Paths.get("src\\main\\resources\\week1\\input.sql")).mapToInt(Integer::valueOf).sum());
    }
}
