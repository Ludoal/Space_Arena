components {
  id: "star"
  component: "/prototypes/star.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"dot_soft_128\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/celestial.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.1
    y: 0.1
    z: 0.1
  }
}
