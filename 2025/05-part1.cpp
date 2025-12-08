#include <bits/stdc++.h>

#include <algorithm>
#include <boost/algorithm/string.hpp>
#include <fstream>
#include <ios>
#include <iostream>
#include <list>
#include <utility>
#include <vector>

int main(int, char** argv) {
  auto part = std::stoi(argv[1]);
  auto filename = argv[2];
  int result = 0;
  std::ifstream file(filename, std::ios_base::in);
  std::list<std::pair<long, long>> ingredients;
  std::string line;
  while (getline(file, line)) {
    if (line == "") {
      break;
    }
    std::vector<std::string> range;
    boost::split(range, line, boost::is_any_of("-"));
    ingredients.emplace_back(std::stol(range[0]), std::stol(range[1]));
  }

  while (getline(file, line)) {
    auto value = std::stol(line);
    for (auto range : ingredients) {
      if (value >= range.first && value <= range.second) {
        result++;
        std::cerr << value << std::endl;
        break;
      }
    }
  }
  std::cout << result << std::endl;
}
