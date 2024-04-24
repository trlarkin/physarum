import numpy as np
import matplotlib.pyplot as plt
from PIL import Image
import sys


def display_image_with_log_intensity(filepath):
    # Read the image
    img = Image.open(filepath)

    # Convert the image to grayscale
    img_gray = img.convert('L')

    # Convert the image to a numpy array
    img_array = np.array(img_gray)

    # Apply logarithmic transformation to intensity
    img_log = np.log1p(img_array)

    # Display the image with log intensity using matplotlib
    fig, axs = plt.subplots(1, 2)
    axs: list[plt.Axes]

    axs[0].imshow(img_array, cmap='gray')
    axs[0].set_title(f"{filepath} original")
    axs[0].axis('off')

    axs[1].imshow(img_log, cmap='gray')
    axs[1].set_title(f"{filepath} logified")
    axs[1].axis('off')

    plt.show(block=False)


# Loop over all input filenames
for filepath in sys.argv[1:]:
    display_image_with_log_intensity(filepath)
input("Enter to close ...")
