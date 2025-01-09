components {
  id: "asteroid"
  component: "/prototypes/asteroids/asteroid.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"a8-1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/tilesources/asteroid8.tilesource\"\n"
  "}\n"
  ""
  scale {
    x: 0.5
    y: 0.5
    z: 0.5
  }
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"asteroids\"\n"
  "mask: \"asteroids\"\n"
  "mask: \"border\"\n"
  "mask: \"bullets_enemy\"\n"
  "mask: \"bullets_ship\"\n"
  "mask: \"ship\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 47.542374\n"
  "  data: 50.01775\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
