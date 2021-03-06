## mackerel-plugin-linux-network

This plugin collects network-related statistics from Linux kernel.

### Requirements
Linux needed. Older kernels support fewer statistics.

### Usage
Just run `mackerel-plugin-linux-network`.

In `mackerel-agent.conf`:
```toml
[plugin.metrics.network]
command="/usr/local/bin/mackerel-plugin-linux-network"
```

#### General options
```
-prefix    a prefix string to the metrics names (defaule: "network.")
-tempfile  a path to temporary file in which metrics values are stored in order to calucurate difference
```

#### Filtering options
The command displays all the supported metrics by default (more than 100!). By using filtering options you can filter out some of them as needed.
Check out `mackerel-plugin-linux-network -help` for all the available options.

Here are some examples:
```
-all           Enable all metrics
-all=true      Same as -all
-all=false     Disable all metrics
-ip            Enable all metrics under `ip.*`
-ip=false      Disable all metrics under `ip.*`
-ip-datagrams  Enable all metrics undler `ip.datagrams.*`
```

Note that the option given later in the command line takes the precedence:
```
mackerel-plugin-linux-network -all=false -ip=true  # Show only IP-related metrics
mackerel-plugin-linux-network -all=true -ip=false  # Show everything but IP-related metrics
mackerel-plugin-linux-network -ip=false -all=true  # Show everything
mackerel-plugin-linux-network -ip=true  -all=false # Show nothing
```

## License

```
Copyright 2016 Kasumi Hanazuki

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
