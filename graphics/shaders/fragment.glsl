uniform vec4 color;
uniform sampler2D sampler_obj;

void main() {
    output_color = color * texture(sampler_obj, tex_coords.xy);
}