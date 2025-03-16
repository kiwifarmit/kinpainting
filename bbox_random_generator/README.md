To install and use the node:

1. Create a `custom_nodes` folder in the ComfyUI main directory if it doesn't already exist.

2. Inside `custom_nodes`, create a new folder for the node, for example, `rectangle_mask`.

3. In the `rectangle_mask` folder, create two files:
    - `__init__.py` (empty)
    - `rectangle_mask_node.py` (with the node code)

4. Restart ComfyUI.

To use it in the workflow:

1. The node will appear as "Rectangle Mask Generator" in the "mask" category.

2. Connect it to an inpainting node as an input mask.

3. Configure the parameters:
    - Image dimensions (width/height)
    - Position (center_x/y) and size (rect_width/height) of the rectangle
    - Activate `randomize` for random values
    - Activate `save_output` to save the mask and data

**Example 1:  Creating a centered square mask**

1. **Node Configuration:**
    *  **width:** 512
    *  **height:** 512
    *  **center_x:** 256 
    *  **center_y:** 256
    *  **rect_width:** 128
    *  **rect_height:** 128
    *  **randomize:** unchecked
    *  **save_output:** unchecked

2. **Workflow Connection:**
    * Connect the "Mask" output of the Rectangle Mask Generator to the "Mask" input of an inpainting node (e.g., "SDXL Inpaint").

This configuration will generate a 128x128 pixel square mask in the center of a 512x512 image. This mask can then be used to inpaint a specific area of the image using the connected inpainting node.

**Example 2: Randomizing mask position and size**

1. **Node Configuration:**
    *  **width:** 512
    *  **height:** 512
    *  **center_x:** 256 
    *  **center_y:** 256
    *  **rect_width:** 128
    *  **rect_height:** 128
    *  **randomize:** checked
    *  **save_output:** checked

2. **Workflow Connection:**
    *  Connect the "Mask" output to the "Mask" input of an inpainting node.

In this case, with `randomize` enabled, the node will generate a mask with a random position and size within the specified range (centered around `center_x` and `center_y`). Enabling `save_output` will save the generated mask as an image file and the mask data as a JSON file, which can be useful for debugging or reusing the mask later.

**Example 3: Using the mask for outpainting**

You can also use the Rectangle Mask Generator for outpainting.  Instead of connecting it to an inpainting node, connect it to the "Mask" input of an outpainting node (e.g., "SDXL Outpaint"). This will allow you to generate content outside the original image boundaries, within the area defined by the rectangular mask.

**Tips:**

* Experiment with different parameter values to achieve various masking effects.
* Use the `randomize` option to create a variety of masks and explore different inpainting/outpainting results.
*  Visualize the generated mask by connecting the "Mask" output to a "Save Image" node to ensure it matches your desired area of modification.

