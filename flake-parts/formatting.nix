{
  inputs,
  self,
  ...
}: {
  perSystem = {
    pkgs,
    lib,
    self',
    ...
  }: let
    formatters = [
      pkgs.alejandra
      self'.packages.go
    ];

    treefmt = pkgs.writeShellApplication {
      name = "treefmt";
      runtimeInputs = [pkgs.treefmt] ++ formatters;
      text = ''
        exec treefmt "$@"
      '';
    };
  in {
    packages = {
      inherit treefmt;
    };

    legacyPackages = {
      inherit formatters;
    };
  };
}
