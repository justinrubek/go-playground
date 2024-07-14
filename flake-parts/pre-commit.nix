{inputs, ...}: {
  imports = [
    inputs.pre-commit-hooks.flakeModule
    ./formatting.nix
  ];
  perSystem = {self', ...}: {
    pre-commit = {
      check.enable = true;

      settings = {
        src = ../.;
        hooks = {
          statix.enable = true;
          treefmt = {
            enable = true;
            package = self'.packages.treefmt;
          };
        };
      };
    };
  };
}
