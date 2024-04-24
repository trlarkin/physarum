import os
import sys
import imageio


def create_gif(directory: str):
    # Get list of PNG files in the directory
    png_files = [filename for filename in os.listdir(
        directory) if filename.endswith('.png')]

    # Sort the PNG files
    png_files.sort()

    images = []
    for filename in png_files:
        filepath = os.path.join(directory, filename)
        # Read each PNG image and append to the list
        images.append(imageio.imread(filepath))

    # Save the images as a GIF animation
    # output_file = os.path.join(directory, f'../test.gif')
    output_file = os.path.join(directory, f'../animation.gif')
    imageio.mimsave(output_file, images)

    print(f"GIF animation saved as '{output_file}'")


if __name__ == "__main__":
    # Input directory name
    directory_name = sys.argv[1]

    # Call the function to create the GIF
    create_gif(directory_name)
