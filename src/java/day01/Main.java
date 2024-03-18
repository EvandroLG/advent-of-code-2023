import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Map;
import java.util.Optional;

public class Main {
  private static final Map<String, Integer> map = Map.of("one", 1, "two", 2, "three", 3, "four", 4,
      "five", 5, "six", 6, "seven", 7, "eight", 8, "nine", 9);

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
      final var digits = getDigits(line);
      final var first = digits.stream().findFirst();
      final var last = digits.stream().reduce((f, second) -> second);

      return concatenateAndConvertToInt(first, last);
    }).sum();
  }

  private static ArrayList<Integer> getDigits(String line) {
    final var result = new ArrayList<Integer>();
    int i = 0;

    while (i < line.length()) {
      final var c = line.charAt(i);

      if (Character.isDigit(c)) {
        result.add(Character.getNumericValue(c));
      } else {
        final var word = getWordInMap(line, i);

        if (word != null) {
          result.add(map.get(word));
        }
      }

      i++;
    }

    return result;
  }

  private static String getWordInMap(String line, int startIndex) {
    for (int i = startIndex + 1; i <= line.length(); i++) {
      final var word = line.substring(startIndex, i);
      if (map.containsKey(word)) {
        return word;
      }
    }

    return null;
  }

  private static int concatenateAndConvertToInt(Optional<Integer> first, Optional<Integer> last) {
    final var result = new StringBuilder();
    first.ifPresent(result::append);
    last.ifPresent(result::append);

    return result.isEmpty() ? 0 : Integer.parseInt(result.toString());
  }
}
