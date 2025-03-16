## Information about the work done with the ComfyUI tool.

For image generation done with Ideogram, please refer to the readme in the "ideogram-tool" folder.

## Tools

  * ComfyUI: [https://github.com/comfyanonymous/ComfyUI](https://www.google.com/url?sa=E&source=gmail&q=https://github.com/comfyanonymous/ComfyUI)
  * ComfyUI Manager: [https://github.com/ltdrdata/ComfyUI-Manager](https://www.google.com/url?sa=E&source=gmail&q=https://github.com/ltdrdata/ComfyUI-Manager)
  * ComfyUI KJNodes: [https://github.com/kijai/ComfyUI-KJNodes](https://www.google.com/url?sa=E&source=gmail&q=https://github.com/kijai/ComfyUI-KJNodes)
  * ComfyUI Impact Pack: [https://github.com/ltdrdata/ComfyUI-Impact-Pack](https://www.google.com/url?sa=E&source=gmail&q=https://github.com/ltdrdata/ComfyUI-Impact-Pack)

Our custom node is in the bbox\_random\_generator folder.

## Model used:

  * epiCRealism\_naturalSinRC1VAE: [https://civitai.com/models/25694?modelVersionId=143906](https://www.google.com/url?sa=E&source=gmail&q=https://civitai.com/models/25694?modelVersionId=143906)
    This should be placed in the "models\\checkpoints" folder under ComfyUI.

## ComfyUI Scripts

The ComfyUI scripts are in the comfyui-workflows\\ folder.

The nodes are grouped for readability. In particular:

  * Group 'Create "without feature" image': generates images without the feature of interest.
  * Group 'Create "with feature" images': generates images with the feature of interest.
  * Group 'Create "with feature" by Inpainting': adds the "feature" to "without feature" images.

## Generating images with ComfyUI

1.  Install all the tools.
2.  Start with run\_nvidia\_gpu.bat.
3.  Use the Load button to load the script you are interested in.
4.  Press the Queue Prompt button to generate one image with and one image without target object.