
void main() {
    gl_Position = mvp * vec4(in_position, 1.0);
    tex_coords = in_tex_coords;
}