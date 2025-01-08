components {
  id: "ship"
  component: "/prototypes/ship.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"DKO-api-X1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/ships.atlas\"\n"
  "}\n"
  ""
  position {
    y: 20.0
  }
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
  "group: \"ship\"\n"
  "mask: \"asteroids\"\n"
  "mask: \"border\"\n"
  "mask: \"powerups\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      y: 14.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 38.652122\n"
  "}\n"
  "linear_damping: 1.0\n"
  "angular_damping: 0.9\n"
  ""
}
embedded_components {
  id: "bullet"
  type: "factory"
  data: "prototype: \"/prototypes/bullet.go\"\n"
  ""
}
embedded_components {
  id: "exhaustFwd"
  type: "sprite"
  data: "default_animation: \"exhaust\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/ships.atlas\"\n"
  "}\n"
  ""
  position {
    y: -47.0
    z: -1.0
  }
  rotation {
    z: 1.0
    w: 6.123234E-17
  }
  scale {
    x: 0.8
    y: 1.5
  }
}
embedded_components {
  id: "exhaustBwd2"
  type: "sprite"
  data: "default_animation: \"exhaust\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 16.0\n"
  "  y: 41.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/ships.atlas\"\n"
  "}\n"
  ""
  position {
    x: -32.0
    y: 12.0
    z: -1.0
  }
  scale {
    x: 0.5
    y: 0.8
  }
}
embedded_components {
  id: "exhaustLatL"
  type: "sprite"
  data: "default_animation: \"exhaust\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 16.0\n"
  "  y: 41.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/ships.atlas\"\n"
  "}\n"
  ""
  position {
    x: 28.0
    y: 26.0
    z: -1.0
  }
  rotation {
    z: -0.70710677
    w: 0.70710677
  }
  scale {
    x: 0.5
    y: 0.8
  }
}
embedded_components {
  id: "exhaustLatR"
  type: "sprite"
  data: "default_animation: \"exhaust\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 16.0\n"
  "  y: 41.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/ships.atlas\"\n"
  "}\n"
  ""
  position {
    x: -28.0
    y: 26.0
    z: -1.0
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
  scale {
    x: 0.5
    y: 0.8
  }
}
embedded_components {
  id: "exhaustRotL"
  type: "sprite"
  data: "default_animation: \"exhaust\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 16.0\n"
  "  y: 41.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/ships.atlas\"\n"
  "}\n"
  ""
  position {
    x: 20.0
    y: 50.0
    z: -1.0
  }
  rotation {
    z: -0.5
    w: 0.8660254
  }
  scale {
    x: 0.5
    y: 0.8
  }
}
embedded_components {
  id: "exhaustRotR"
  type: "sprite"
  data: "default_animation: \"exhaust\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 16.0\n"
  "  y: 41.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/ships.atlas\"\n"
  "}\n"
  ""
  position {
    x: -20.0
    y: 50.0
    z: -1.0
  }
  rotation {
    z: 0.5
    w: 0.8660254
  }
  scale {
    x: 0.5
    y: 0.8
  }
}
embedded_components {
  id: "exhaustBwd1"
  type: "sprite"
  data: "default_animation: \"exhaust\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 16.0\n"
  "  y: 41.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/ships.atlas\"\n"
  "}\n"
  ""
  position {
    x: 32.0
    y: 12.0
    z: -1.0
  }
  scale {
    x: 0.5
    y: 0.8
  }
}
