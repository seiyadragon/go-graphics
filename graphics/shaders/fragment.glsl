#version 410 core

in vec3 tex_coords;
out vec4 output_color;
in float tex;

uniform vec4 color;
uniform sampler2D sampler_obj;
uniform sampler2D sampler_obj2;

void main() {
    if (tex == 1.0)
        output_color = color * texture(sampler_obj, tex_coords.xy);
    else
        output_color = color * texture(sampler_obj2, tex_coords.xy);
}