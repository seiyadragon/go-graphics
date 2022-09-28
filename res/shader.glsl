#version 410 core

layout (location = 0) in vec3 in_position;
layout (location = 1) in vec3 in_tex_coords;
layout (location = 2) in vec3 in_normal;
out vec3 tex_coords;

uniform mat4 mvp;

void main() {
    gl_Position = mvp * vec4(in_position, 1.0);
    tex_coords = in_tex_coords;
}

{separator}
#version 410 core

in vec3 tex_coords;
out vec4 output_color;

uniform sampler2D sampler_obj;

void main() {
        output_color = texture(sampler_obj, tex_coords.xy);
}