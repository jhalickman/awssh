awssh
=====

awssh allows you to easily SSH into your running EC2 instances by name with out having to remember the private IP or which key you need to use.


## Installation

#### Via Go

```bash
$ go get github.com/jhalickman/awssh
```

#### Binaries

- **mac** [386]() / [amd64]()


## Configuration
awss uses [viper](http://www.github.com/spf13/viper) for config so you can specify your config as a JSON, TOML or YAML file. The file must be named `config.[json|toml|yaml]` and placed in either `$HOME/.awssh` or the location of the binary.

#### Config values
<table>
  <tr> 
	  <th>Config Name</th>
	  <th>required (default)</th>
	  <th>description</th>
  </tr>
  <tr>
  	<td>access_token</td>
  	<td>Y</td>
  	<td>AWS account access token</td>
  </tr>
  <tr>
  	<td>access_secret</td>
  	<td>Y</td>
  	<td>AWS account secret</td>
  </tr>
  <tr>
  	<td>key_folder</td>
  	<td> N ($HOME/.awssh/keys)</td>
  	<td>Directory where your AWS SSH keys are stored</td>
  </tr>
  <tr>
  	<td>login_name</td>
  	<td>N (ubuntu) </td>
  	<td>login name used when connecting via SSH</td>
  </tr>
</table>

## Usage

#### select instance running instances from a list by name
```bash
$ awssh
0)web-01
1)web-02
3)app-01
4)db-01
Which server do you want to login to? [Enter 0-4 or instance name]
```

####specify instance name in command
```bash
$ awssh web-02
```