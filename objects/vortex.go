components {
  id: "vortex"
  component: "/prototypes/vortex.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/tilesources/vortex.tilesource\"\n"
  "}\n"
  ""
  scale {
    x: 4.0
    y: 4.0
    z: 4.0
  }
}
