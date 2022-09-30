#ifndef CONFIG_H
#define CONFIG_H

#include <string>

typedef struct config {
  std::string endpoint;
} config_t;

int config_parse(int argc, char** argv, config_t& config);

#endif  // CONFIG_H
