{inputs, ...}: {
  perSystem = {
    config,
    pkgs,
    system,
    inputs',
    self',
    lib,
    ...
  }: let
    inherit (self'.packages) go golangci-lint treefmt;

    devTools = [
      go
      golangci-lint
      treefmt
    ];
  in {
    devShells = {
      default = pkgs.mkShell rec {
        packages = devTools;

        shellHook = ''
          ${config.pre-commit.installationScript}
        '';
      };
    };
  };
}
