embedded_components {
  id: "ground1"
  type: "sprite"
  data: "default_animation: \"ground\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/game.atlas\"\n"
  "}\n"
  ""
  position {
    y: 6.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"ground\"\n"
  "mask: \"hero\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: 15.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 450.0\n"
  "  data: 15.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "signleft"
  type: "sprite"
  data: "default_animation: \"signLeft\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/game.atlas\"\n"
  "}\n"
  ""
  position {
    x: 354.0
    y: 60.0
  }
}
embedded_components {
  id: "ground2"
  type: "sprite"
  data: "default_animation: \"ground\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/game.atlas\"\n"
  "}\n"
  ""
  position {
    x: -800.0
    y: 6.0
  }
}
embedded_components {
  id: "ground3"
  type: "sprite"
  data: "default_animation: \"ground\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/game.atlas\"\n"
  "}\n"
  ""
  position {
    x: 800.0
    y: 6.0
  }
}
embedded_components {
  id: "signright"
  type: "sprite"
  data: "default_animation: \"signRight\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/game.atlas\"\n"
  "}\n"
  ""
  position {
    x: -354.0
    y: 60.0
  }
}
