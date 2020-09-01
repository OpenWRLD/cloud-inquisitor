locals {
    build_files = fileset("${var.working_dir}/cloud-inquisitor/builds/*")
}

resource "null_resource" "local_machine" {

    provisioner "local-exec" {
        working_dir = var.working_dir
        command = <<CMD
            git clone -b ${var.branch} https://github.com/RiotGames/cloud-inquisitor.git &&\
            cd cloud-inquisitor &&\
            make build
        CMD

        environment = {
            SETTINGS_FILE = var.settings_file
        }
    }
    
    triggers = {
        builds = build_files
    }
}

variable "branch" {
    type = string
    description = "branch of the project to clone and build from"
    default = "master"
}

variable "working_dir" {
    type = string
    description = "path to clone repo into"
    default = "./"
}

variable "settings_file" {
    type = string
    description = "path to settings file"
}

output "binary_path" {
    value = "${var.working_dir}/cloud-inquisitor/builds"
    depends_on = [
        null_resource.local_machine
    ]
}