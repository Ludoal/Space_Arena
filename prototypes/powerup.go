components {
  id: "powerup"
  component: "/prototypes/powerup.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"ammo\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/tilesources/powerups.tilesource\"\n"
  "}\n"
  ""
  position {
    z: -1.0
  }
  scale {
    x: 0.5
    y: 0.5
  }
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"powerups\"\n"
  "mask: \"ship\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      x: 3.0\n"
  "      y: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 62.0\n"
  "}\n"
  ""
}
