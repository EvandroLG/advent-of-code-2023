import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Optional;
import java.util.stream.Collector;
import java.util.stream.Collectors;

public class Main {
  public static void main(String[] args) {
    final var lines = readFromInputFile("./input.txt");
    final var result = calculateCalibration(lines);
    System.out.println("Result: " + result);
  }

  private static ArrayList<String> readFromInputFile(String filepath) {
    final var lines = new ArrayList<String>();

    try (final var br = new BufferedReader(new FileReader(filepath))) {
      String line;

      while ((line = br.readLine()) != null) {
        lines.add(line);
      }
    } catch (IOException e) {
      System.err.println(e);
    }

    return lines;
  }

  private static Integer calculateCalibration(ArrayList<String> lines) {
    return lines.stream().mapToInt(line -> {
      final var digits = line.chars().filter(Character::isDigit)
          .mapToObj(Character::getNumericValue).collect(Collectors.toList());
      final var first = digits.stream().findFirst();
      final var last = digits.stream().reduce((f, second) -> second);

      return concatenateAndConvertToInt(first, last);
    }).sum();
  }

  private static int concatenateAndConvertToInt(Optional<Integer> first, Optional<Integer> last) {
    final var result = new StringBuilder();
    first.ifPresent(result::append);
    last.ifPresent(result::append);

    return Integer.parseInt(result.toString());
  }
}
