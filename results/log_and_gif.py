import os
import sys
import numpy as np
import matplotlib.pyplot as plt
from PIL import Image
import imageio


def convert_to_log_intensity(image):
    # Convert the image to grayscale
    img_gray = image.convert('L')

    # Convert the image to a numpy array
    img_array = np.array(img_gray)

    # Apply logarithmic transformation to intensity
    img_log = np.log1p(img_array)

    return img_log


def process_images_in_directory(directory):
    # Create a list to store log intensity versions of images
    log_images = []

    # Loop over each file in the directory
    for filename in os.listdir(directory):
        filepath = os.path.join(directory, filename)
        # Check if the file is an image
        if os.path.isfile(filepath) and filename.lower().endswith(('.png', '.jpg', '.jpeg')):
            # Read the image
            img = Image.open(filepath)
            # Convert to log intensity
            img_log = convert_to_log_intensity(img)
            log_images.append(img_log)

    return log_images


def create_gif(images, output_filename='output.gif', duration=0.2):
    # Save the images as frames of a GIF
    with imageio.get_writer(output_filename, mode='I', duration=duration) as writer:
        for image in images:
            # Convert to RGB format (required by imageio)
            rgb_image = np.uint8(plt.cm.gray(image) * 255)
            # Append image to the GIF
            writer.append_data(rgb_image)


def main(directory):
    # Process images in the directory
    log_images = process_images_in_directory(directory)

    # Create a GIF from the log intensity versions of the images
    create_gif(log_images, output_filename='log_intensity.gif', duration=2)

    print("GIF created successfully!")


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py directory")
        sys.exit(1)
    directory = sys.argv[1]
    main(directory)
