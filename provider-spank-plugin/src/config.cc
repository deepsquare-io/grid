#include "config.h"

extern "C" {
#include <slurm/spank.h>

extern const char plugin_type[];
}

#include <string>

#include "report.h"

int config_parse(int argc, char **argv, config_t &config) {
  config.endpoint = "localhost:3000";

  for (int i = 0; i < argc; i++) {
    if (std::string("endpoint=").compare(0, 9, argv[i])) {
      config.endpoint = std::string(argv[i] + 9);
    } else {
      slurm_error("%s: unknown configuration option: %s", plugin_type, argv[i]);
      return 1;
    }
  }

  return 0;
}
