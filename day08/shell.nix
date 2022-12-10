with import <nixpkgs> {};
let
  pythonEnv = python310.withPackages (ps: [
    ps.numpy
    ps.ipython
    ps.pytest
  ]);
in mkShell {
  packages = [
     pythonEnv
  ];
}
