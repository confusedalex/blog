{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    systems.url = "github:nix-systems/default";
  };

  outputs =
    { self, systems, nixpkgs, ... }@inputs:
    let
      eachSystem = f: nixpkgs.lib.genAttrs (import systems) (system: f nixpkgs.legacyPackages.${system});
    in
    {
      packages = eachSystem (pkgs: {
        fetch-webmentions = pkgs.writeShellApplication {
          name = "fetch-webmentions";
          runtimeInputs = with pkgs; [
            curl
            jq
          ];
          text = (builtins.readFile ./fetch-webmentions.sh);
        };
      });

      devShells = eachSystem (pkgs: {
        default = pkgs.mkShell {
          buildInputs = with pkgs; [
            zola
            jq
            self.packages.${system}.fetch-webmentions
          ];
        };
      });
    };
}
