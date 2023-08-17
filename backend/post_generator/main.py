# from PIL import Image, ImageDraw, ImageFont
#
# def generate_instagram_post(text):
#     # Image dimensions (1080x1080 is a common Instagram post size)
#     width, height = 1080, 1080
#     
#     # Create a blank white image
#     img = Image.new('RGB', (width, height), color='white')
#     
#     # Define drawing context
#     d = ImageDraw.Draw(img)
#     
#     # Define font and size
#     # font = ImageFont.truetype("arial.ttf", 50)  # Change the font path and size as required
#     font = ImageFont.load_default()
#     
#     # Calculate text width and height to center the text
#     # text_width, text_height = d.textsize(text, font=font)
#     # text_width, text_height = d.text(xy, text)
#     text_width = 50 
#     text_height = 50
#     x = (width - text_width) / 2
#     y = (height - text_height) / 2
#
#     # Draw the text on the image
#     d.text((x, y), text, fill='black', font=font)
#     
#     # Save the image
#     img.save('instagram_post.png')
#
# text = "I did a shit on your mum, I did as shit on your mum, I did a shit on your mum, and she rather liked it"
# generate_instagram_post(text)
from PIL import Image, ImageDraw, ImageFont

def generate_instagram_dm(sender: str, message: str):
    # Image dimensions (typical phone screenshot dimensions)
    width, height = 1080, 1920

    # Create a blank image with a light gray background (Instagram DM background color)
    img = Image.new('RGB', (width, height), color='#FAFAFA')

    # Define drawing context
    d = ImageDraw.Draw(img)

    # Load a font (for better representation, use the font Instagram uses or something similar)
    font = ImageFont.load_default()

    # Calculate text width and height
    text_width = 500
    text_height = 500

    # Define padding, bubble size, and positioning
    padding = 20
    bubble_width = text_width + 3*padding  # Additional padding for aesthetics
    bubble_height = text_height + 2*padding
    bubble_x = width - bubble_width - padding
    bubble_y = height - bubble_height - 3*padding  # Positioning it a bit above the bottom
    text_x = bubble_x + padding
    text_y = bubble_y + padding

    # Draw the message bubble (light blue for the sender's message)
    d.rounded_rectangle([bubble_x, bubble_y, bubble_x+bubble_width, bubble_y+bubble_height], radius=20, fill="#E6EBF2")

    # Draw the text on the message bubble
    d.text((text_x, text_y), message, fill='black', font=font)

    # (Optional) You can add more details like sender's name, profile pic, timestamp, etc.

    # Save the image
    img.save('instagram_dm.png')

# Example usage:
generate_instagram_dm("SenderName", "Hello! This is a mock Instagram DM!")
from PIL import Image, ImageDraw, ImageFont

def generate_instagram_post(text):
    # Image dimensions
    width, height = 1080, 1920  # Typical phone screenshot dimensions

    # Create a gradient background
    img = Image.new('RGB', (width, height), color='#fafafa')

    # Define drawing context
    d = ImageDraw.Draw(img)

    # Define font and size
    # It's recommended to use a font that's closer to Instagram's
    font = ImageFont.load_default()  # or ImageFont.truetype("path_to_font.ttf", 40)

    # Calculate text width and height
    # text_width, text_height = d.multiline_textsize(text, font=font)
    text_width = 50 
    text_height = 50

    # Define message bubble padding and size
    padding = 20
    bubble_width = text_width + 2*padding
    bubble_height = text_height + 2*padding

    # Define the position of the text and message bubble
    x_bubble = 30
    y_bubble = height - bubble_height - 30  # position at the bottom
    x_text = x_bubble + padding
    y_text = y_bubble + padding

    # Draw the message bubble
    d.rounded_rectangle([x_bubble, y_bubble, x_bubble+bubble_width, y_bubble+bubble_height], radius=20, fill="#ECECEC")

    # Draw the text inside the bubble
    d.multiline_text((x_text, y_text), text, fill='black', font=font)

    # (Optional) Add a profile picture, username, and timestamp here

    # Save the image
    img.save('instagram_dm.png')

text = "I did a shit on your mum, I did as shit on your mum, I did a shit on your mum, and she rather liked it"
generate_instagram_post(text)

