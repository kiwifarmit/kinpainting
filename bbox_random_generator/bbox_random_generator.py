import random
import json
import os
import folder_paths

def get_unique_seed():
    return random.randint(0, 2**64 - 1)

class BboxRandomGenerator:
    @classmethod
    def INPUT_TYPES(cls):
        return {
            "required": {
                "rand_x_min": ("INT", {"default": 0, "min": 0, "max": 2048}),
                "rand_x_max": ("INT", {"default": 512, "min": 0, "max": 2048}),
                "rand_y_min": ("INT", {"default": 0, "min": 0, "max": 2048}),
                "rand_y_max": ("INT", {"default": 512, "min": 0, "max": 2048}),
                "rand_w_min": ("INT", {"default": 64, "min": 1, "max": 2048}),
                "rand_w_max": ("INT", {"default": 256, "min": 1, "max": 2048}),
                "rand_h_min": ("INT", {"default": 64, "min": 1, "max": 2048}),
                "rand_h_max": ("INT", {"default": 256, "min": 1, "max": 2048}),
                "save_data": ("BOOLEAN", {"default": False}),
                "filename_prefix": ("STRING", {"default": "bbox"})
            }
        }
    
    RETURN_TYPES = ("INT", "INT", "INT", "INT")
    RETURN_NAMES = ("location_x", "location_y", "rect_width", "rect_height")
    FUNCTION = "generate_params"
    CATEGORY = "utils"

    def get_latest_counter(self, folder_path, filename_prefix):
        counter = 1
        if not os.path.exists(folder_path):
            return counter
        try:
            files = [f for f in os.listdir(folder_path) if f.startswith(filename_prefix) and f.endswith('.json')]
            if files:
                counters = [int(f[len(filename_prefix)+1:-5]) for f in files if f[len(filename_prefix)+1:-5].isdigit()]
                if counters:
                    counter = max(counters) + 1
        except Exception as e:
            print(f"Error finding latest counter: {e}")
        return counter

    @classmethod
    def IS_CHANGED(cls, **kwargs):
        return float("nan")

    def save_params(self, data, filename_prefix):
        try:
            output_dir = folder_paths.get_output_directory()
            counter = self.get_latest_counter(output_dir, filename_prefix)
            filepath = os.path.join(output_dir, f"{filename_prefix}_{counter:05}.json")
            
            with open(filepath, 'w') as f:
                json.dump(data, f, indent=2)
        except Exception as e:
            print(f"Error saving data: {e}")

    def generate_params(self, rand_x_min, rand_x_max, rand_y_min, rand_y_max,
                       rand_w_min, rand_w_max, rand_h_min, rand_h_max,
                       save_data=False, filename_prefix="bbox"):
        seed = get_unique_seed()
        random.seed(seed)
        
        x = random.randint(rand_x_min, rand_x_max)
        y = random.randint(rand_y_min, rand_y_max)
        w = random.randint(rand_w_min, rand_w_max)
        h = random.randint(rand_h_min, rand_h_max)
        
        if save_data:
            params_data = {
                "location_x": x,
                "location_y": y,
                "rect_width": w,
                "rect_height": h,
                "seed": seed
            }
            self.save_params(params_data, filename_prefix)
        
        return (x, y, w, h)

NODE_CLASS_MAPPINGS = {
    "BboxRandomGenerator": BboxRandomGenerator
}

NODE_DISPLAY_NAME_MAPPINGS = {
    "BboxRandomGenerator": "BBox Random Generator"
}