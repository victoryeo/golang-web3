// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Example {
  string public name;
  uint256 public count;
  string public a;
  string public b;

  event Hello(string name, uint256 count);

  function hello(string calldata _name) public {
    name = _name;
    count += 1;
    emit Hello(_name, count);
  }

  function exEncode () external {
    a = string(abi.encode("hello", " world"));
    b = string(abi.encodePacked("hello", " world"));
  }
}

