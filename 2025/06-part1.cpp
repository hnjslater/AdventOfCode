#include <array>
#include <boost/algorithm/string.hpp>
#include <boost/algorithm/string_regex.hpp>
#include <fstream>
#include <ios>
#include <iostream>
#include <set>
#include <string>
#include <utility>
#include <vector>
#include <numeric>
#include <boost/multiprecision/cpp_int.hpp>

int main(int argc, char **argv) {
  auto part = std::stoi(argv[1]);
  auto filename = argv[2];

  std::ifstream file(filename, std::ios_base::in);
  std::string line;
  std::vector<std::vector<int>> numbers;
  boost::multiprecision::cpp_int result = 0;
  while (getline(file, line)) {
    std::vector<std::string> values;
    boost::split_regex(values, line, boost::regex("[[:blank:]]+"));
    if (numbers.size() == 0) {
      numbers.resize(values.size());
    }
    auto j = 0;
    for (auto i = 0; i < values.size(); i++) {
      if (values[i] == "+") {
	      boost::multiprecision::cpp_int r1 = 0;
	      
	      result += std::accumulate(numbers[i].begin(), numbers[i].end(), r1);
      } else if (values[i] == "*") {
	      boost::multiprecision::cpp_int r1 = 1;
	      result += std::accumulate(numbers[i].begin(), numbers[i].end(), r1, std::multiplies<boost::multiprecision::cpp_int>());
      } else if (values[i].size() > 0) {
        numbers[j].push_back(std::stoi(values[i]));
	j++;
      }
    }
  }


  std::cout << result << std::endl;
}
