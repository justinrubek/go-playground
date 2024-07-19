{inputs, ...}: {
  perSystem = {
    config,
    pkgs,
    system,
    inputs',
    self',
    lib,
    ...
  }: {
    packages = {
      inherit (inputs'.nix-go.packages) go golangci-lint;
    };
  };
}
