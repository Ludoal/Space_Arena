components {
  id: "enemy"
  component: "/objects/enemy.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"WO84-wu-X1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/ships.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.5
    y: 0.5
  }
}
