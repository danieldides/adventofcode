package week1;

import java.io.IOException;
import java.lang.reflect.Array;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import java.util.stream.Stream;

public class Application {

    private static final String PATH_TO_INPUT = "src/main/resources/week1/input.sql";

    private static Stream<String> getLines() throws IOException {
        return Files.lines(Paths.get(PATH_TO_INPUT));
    }

    private static int problemOne() throws IOException {
        return getLines().mapToInt(Integer::valueOf).sum();
    }

    private static int problemTwo() throws IOException {
        int freq = 0;
        List<Integer> inputs  = getLines().mapToInt(Integer::valueOf).boxed().collect(Collectors.toList());
        ArrayList<Integer> freqsSoFar = new ArrayList();

        freqsSoFar.add(0);

        boolean found = false;
        int idx = 0;
        while(!found) {
            int newFreq = inputs.get(idx) + freqsSoFar.get(freqsSoFar.size() -1);

            if (freqsSoFar.contains(newFreq)) {
                found = true;
                freq = newFreq;
            }
            freqsSoFar.add(newFreq);

            idx++;
            if (idx >= inputs.size()) {
                idx = 0;
            }
        }
        return freq;

    }


    public static void main(String[] args) throws IOException {
        System.out.println("Problem 1: " + problemOne());
        System.out.println("Problem 2: " + problemTwo());
    }
}
