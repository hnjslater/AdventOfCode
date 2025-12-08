#include <bits/stdc++.h>

#include <algorithm>
#include <boost/algorithm/string.hpp>
#include <boost/multiprecision/cpp_int.hpp>
#include <fstream>
#include <ios>
#include <iostream>
#include <list>
#include <utility>
#include <vector>

auto normalise(auto ingredients) {
  auto it = ingredients.begin();
  auto depth = 0;
  while (it != ingredients.end()) {
    auto second = it->second;
    if (depth == 0) {
      it++;
    } else if (!it->second && depth == 1) {
      it++;
    } else {
      auto old = it;
      it++;
      ingredients.erase(old);
    }
    if (second) {
      depth++;
    } else {
      depth--;
    }
  }
  return ingredients;
}

int main(int, char** argv) {
  auto part = std::stoi(argv[1]);
  auto filename = argv[2];
  boost::multiprecision::cpp_int result = 0;
  std::ifstream file(filename, std::ios_base::in);
  std::list<std::pair<long, bool>> ingredients;
  std::string line;
  while (getline(file, line)) {
    if (line == "") {
      break;
    }
    std::vector<std::string> range;
    boost::split(range, line, boost::is_any_of("-"));
    ingredients.emplace_back(std::stol(range[0]), true);
    ingredients.emplace_back(std::stol(range[1]) + 1, false);
  }

  ingredients.sort();

  ingredients = normalise(ingredients);

  for (auto& [x, y] : ingredients) {
    if (y) {
      result -= x;
    } else {
      result += x;
    }
  }
  std::cerr << std::endl;
  std::cout << result << std::endl;
}
