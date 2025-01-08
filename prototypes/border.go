components {
  id: "border"
  component: "/prototypes/border.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"border\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/misc.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 1.0\n"
  "group: \"border\"\n"
  "mask: \"ship\"\n"
  "mask: \"asteroids\"\n"
  "mask: \"bullets\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: 29.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 50.0\n"
  "  data: 25.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
