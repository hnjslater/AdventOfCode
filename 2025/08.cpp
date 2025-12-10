#include <algorithm>
#include <array>
#include <boost/algorithm/string.hpp>
#include <boost/multiprecision/cpp_int.hpp>
#include <fstream>
#include <ios>
#include <iostream>
#include <map>
#include <string>
#include <utility>
#include <vector>

using boost::multiprecision::cpp_int;

using box = std::tuple<int, int, int>;

int main(int argc, char **argv) {
  auto part = std::stoi(argv[1]);
  auto filename = std::string(argv[2]);
  cpp_int result = 0;
  std::ifstream file(filename, std::ios_base::in);
  std::string line;
  auto prev_line = line;
  std::vector<box> boxes;
  while (getline(file, line)) {
    std::vector<std::string> range;
    boost::split(range, line, boost::is_any_of(","));
    box b;
    boxes.emplace_back(std::stoi(range[0]), std::stoi(range[1]),
                       std::stoi(range[2]));
  }

  const size_t WIRES = [&]() {
    if (part == 1) {
      if (filename == "test08.txt") {
        return 10uz;
      } else {
        return 1000uz;
      }
    } else {
      return boxes.size() * boxes.size();
    }
  }();

  std::vector<std::tuple<cpp_int, box, box>> distances;
  for (auto b1 : boxes) {
    for (auto b2 : boxes) {
      if (b1 == b2 || b1 > b2) {
        continue;
      }
      cpp_int distance = 0;
      distance += cpp_int{std::pow(std::get<0>(b1) - std::get<0>(b2), 2)};
      distance += cpp_int{std::pow(std::get<1>(b1) - std::get<1>(b2), 2)};
      distance += cpp_int{std::pow(std::get<2>(b1) - std::get<2>(b2), 2)};

      distances.emplace_back(distance, b1, b2);
    }
  }

  std::sort(distances.begin(), distances.end());

  std::map<box, int> circuits;
  auto circuit_index = 0;
  std::map<int, int> equal_circuits;
  auto circuit_count = 0;
  auto connected_boxes = 0;

  for (size_t i = 0; i < WIRES; i++) {
    auto [d, b1, b2] = distances[i];
    auto b1circuit = circuits[b1];
    auto b2circuit = circuits[b2];

    if (b1circuit == 0 and b2circuit == 0) {
      circuit_index++;
      circuit_count++;
      connected_boxes += 2;
      circuits[b1] = circuit_index;
      circuits[b2] = circuit_index;
    } else if (b1circuit == 0) {
      circuits[b1] = b2circuit;
      connected_boxes++;
    } else if (b2circuit == 0) {
      circuits[b2] = b1circuit;
      connected_boxes++;
    } else if (b1circuit != b2circuit) {
      circuit_count--;
      for (auto it = circuits.begin(); it != circuits.end(); it++) {
        if (it->second == b2circuit) {
          it->second = b1circuit;
        }
      }
    }
    if (circuit_count == 1 and connected_boxes == boxes.size()) {
      std::cerr << std::get<0>(b1) * std::get<0>(b2) << std::endl;
      break;
    }
  }
  if (part == 1) {
    std::map<int, int> circuit_sizes;

    for (auto &[b, c] : circuits) {
      circuit_sizes[c]++;
    }

    std::vector<int> sizes;
    for (auto &[c, s] : circuit_sizes) {
      sizes.push_back(s);
    }
    std::remove(sizes.begin(), sizes.end(), 0);
    std::sort(sizes.begin(), sizes.end());

    result = 1;
    for (int i = 0; i < 3; i++) {
      result *= sizes[sizes.size() - 1 - i];
    }
    std::cerr << result << std::endl;
  }
}
