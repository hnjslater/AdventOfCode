#include <boost/multiprecision/cpp_int.hpp>
#include <fstream>
#include <ios>
#include <iostream>
#include <numeric>
#include <string>
#include <vector>

int main(int argc, char **argv) {
  auto part = std::stoi(argv[1]);
  auto filename = argv[2];

  std::ifstream file(filename, std::ios_base::in);
  std::string line;
  std::vector<std::string> lines;
  boost::multiprecision::cpp_int result = 0;
  while (getline(file, line)) {
    lines.push_back(line);
  }
  bool first = true;
  char op = ' ';
  std::vector<int> numbers;
  for (auto col = 0; col < lines[0].size(); col++) {
    if (first) {
      op = lines[lines.size() - 1][col];
      first = false;
    }
    std::string number;
    for (auto row = 0; row < lines.size() - 1; row++) {
      if (lines[row][col] != ' ') {
        number += lines[row][col];
      }
    }
    if (number != "") {
      numbers.push_back(std::stoi(number));
    }
    if (number == "" || col + 1 == lines[0].size()) {
      if (op == '+') {
        boost::multiprecision::cpp_int r1 = 0;
        result += std::accumulate(numbers.begin(), numbers.end(), r1);
      } else if (op == '*') {
        boost::multiprecision::cpp_int r1 = 1;
        result +=
            std::accumulate(numbers.begin(), numbers.end(), r1,
                            std::multiplies<boost::multiprecision::cpp_int>());
      }
      numbers.clear();
      first = true;
    }
  }

  std::cout << result << std::endl;
}
