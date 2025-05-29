import os
import shutil  

folder = "***"
destination = {
    "Images": [".png", ".jpg", ".jpeg", ".gif"],
    "Documents": [".pdf", ".docx", ".txt",".doc"],
    "Archives": [".zip", ".rar", ".tar"],
}  

for file in os.listdir(folder):  
    file_path = os.path.join(folder, file)  
    if os.path.isfile(file_path):  
        for folder, extensions in destination.items():  
            if file.endswith(tuple(extensions)):  
                folder_path = os.path.join(folder, folder)  
                os.makedirs(folder_path, exist_ok=True)  
                shutil.move(file_path, folder_path)  

print("folder sorted!")