package client

var schemaStr = `
{
  "description": "The collectorset specification schema",
  "type": "object",
  "required": [
    "spec"
  ],
  "properties": {
    "spec": {
      "type": "object",
      "required": [
        "imageRepository",
        "imageTag",
        "imagePullPolicy",
        "replicas",
        "size",
        "clusterName"
      ],
      "properties": {
        "imageRepository": {
          "description": "The image repository of the collector container",
          "type": "string",
          "default": "logicmonitor/collector"
        },
        "imageTag": {
          "description": "The image tag of the collector container",
          "type": "string",
          "default": "latest"
        },
        "imagePullPolicy": {
          "description": "The image pull policy of the collector container",
          "type": "string",
          "default": "Always",
          "enum": [
            "Always",
            "IfNotPresent",
            "Never"
          ]
        },
        "replicas": {
          "description": "The number of collector replicas",
          "type": "integer",
          "minimum": 1.0,
          "default": 1.0
        },
        "size": {
          "description": "The collector size. Available collector sizes: nano, small, medium, large, extra_large, double_extra_large",
          "type": "string",
          "default": "nano",
          "enum": [
            "nano",
            "small",
            "medium",
            "large",
            "extra_large",
            "double_extra_large"
          ]
        },
        "clusterName": {
          "description": "The clustername of the collector",
          "type": "string"
        },
        "groupID": {
          "description": "The groupId of the collector",
          "type": "integer",
          "minimum": -1.0
        },
        "escalationChainID": {
          "description": "The escalation chain Id of the collectors",
          "type": "integer",
          "minimum": 0.0
        },
        "collectorVersion": {
          "description": "The collector version (Fractional numbered version is invalid. For ex: 29.101 is invalid, correct input is 29101)",
          "type": "integer",
          "minimum": 0.0
        },
        "labels": {
          "description": "The Labels of the collector",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "annotations": {
          "description": "The Annotations of the collector",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "useEA": {
          "description": "Flag to opt for EA collector versions",
          "type": "boolean"
        },
        "proxyURL": {
          "description": "The Http/Https proxy url of the collector",
          "type": "string"
        },
        "secretName": {
          "description": "The Secret resource name of the collector",
          "type": "string"
        },
        "policy": {
          "type": "object",
          "properties": {
            "distributionStrategy": {
              "description": "Distribution strategy to provide collector ID to the client requests from available running collectors",
              "type": "string",
              "default": "RoundRobin"
            },
            "orchestrator": {
              "description": "The container orchestration platform designed to automate the deployment, scaling, and management of containerized applications",
              "type": "string",
              "default": "Kubernetes"
            }
          }
        },
        "statefulsetspec": {
          "x-kubernetes-preserve-unknown-fields": true,
          "description": "The collector StatefulSet specification for customizations",
          "type": "object",
          "properties": {
            "template": {
              "type": "object",
              "properties": {
                "spec": {
                  "type": "object",
                  "properties": {
                    "nodeSelector": {
                      "description": "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
                      "type": "object",
                      "additionalProperties": {
                        "type": "string"
                      }
                    },
                    "priorityClassName": {
                      "description": "If specified, indicates the pod's priority. \"system-node-critical\" and \"system-cluster-critical\" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.",
                      "type": "string"
                    },
                    "containers": {
                      "type": "array",
                      "items": {
                        "description": "A single application container that you want to run within a pod.",
                        "type": "object",
                        "required": [
                          "name"
                        ],
                        "properties": {
                          "name": {
                            "type": "string"
                          },
                          "resources": {
                            "description": "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
                            "type": "object",
                            "properties": {
                              "limits": {
                                "description": "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
                                "type": "object",
                                "additionalProperties": {
                                  "x-kubernetes-int-or-string": true
                                }
                              },
                              "requests": {
                                "description": "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
                                "type": "object",
                                "additionalProperties": {
                                  "x-kubernetes-int-or-string": true
                                }
                              }
                            }
                          }
                        }
                      }
                    },
                    "tolerations": {
                      "type": "array",
                      "items": {
                        "type": "object",
                        "properties": {
                          "effect": {
                            "description": "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
                            "type": "string"
                          },
                          "key": {
                            "description": "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
                            "type": "string"
                          },
                          "operator": {
                            "description": "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
                            "type": "string"
                          },
                          "tolerationSeconds": {
                            "description": "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
                            "format": "int64",
                            "type": "integer"
                          },
                          "value": {
                            "description": "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
                            "type": "string"
                          }
                        }
                      }
                    },
                    "dnsConfig": {
                      "description": "PodDNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.",
                      "type": "object",
                      "properties": {
                        "nameservers": {
                          "description": "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
                          "type": "array",
                          "items": {
                            "type": "string"
                          }
                        },
                        "options": {
                          "description": "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
                          "type": "array",
                          "items": {
                            "description": "PodDNSConfigOption defines DNS resolver options of a pod.",
                            "type": "object",
                            "properties": {
                              "name": {
                                "description": "Required.",
                                "type": "string"
                              },
                              "value": {
                                "type": "string"
                              }
                            }
                          }
                        },
                        "searches": {
                          "description": "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
                          "type": "array",
                          "items": {
                            "type": "string"
                          }
                        }
                      }
                    },
                    "dnsPolicy": {
                      "description": "Set DNS policy for the pod. Defaults to \"ClusterFirst\". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.",
                      "type": "string"
                    },
                    "hostAliases": {
                      "description": "HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified. This is only valid for non-hostNetwork pods.",
                      "type": "array",
                      "items": {
                        "description": "HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.",
                        "type": "object",
                        "properties": {
                          "hostnames": {
                            "description": "Hostnames for the above IP address.",
                            "type": "array",
                            "items": {
                              "type": "string"
                            }
                          },
                          "ip": {
                            "description": "IP address of the host file entry.",
                            "type": "string"
                          }
                        }
                      },
                      "x-kubernetes-patch-merge-key": "ip",
                      "x-kubernetes-patch-strategy": "merge"
                    },
                    "nodeName": {
                      "description": "NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.",
                      "type": "string"
                    },
                    "priority": {
                      "description": "The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.",
                      "type": "integer",
                      "format": "int32"
                    },
                    "restartPolicy": {
                      "description": "Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy",
                      "type": "string"
                    },
                    "schedulerName": {
                      "description": "If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.",
                      "type": "string"
                    },
                    "volumes": {
                      "description": "List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes",
                      "type": "array",
                      "items": {
                        "description": "Volume represents a named volume in a pod that may be accessed by any container in the pod.",
                        "type": "object",
                        "required": [
                          "name"
                        ],
                        "properties": {
                          "name": {
                            "description": "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                            "type": "string"
                          },
                          "awsElasticBlockStore": {
                            "description": "Represents a Persistent Disk resource in AWS.\n\nAn AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.",
                            "type": "object",
                            "required": [
                              "volumeID"
                            ],
                            "properties": {
                              "fsType": {
                                "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
                                "type": "string"
                              },
                              "partition": {
                                "description": "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty).",
                                "type": "integer",
                                "format": "int32"
                              },
                              "readOnly": {
                                "description": "Specify \"true\" to force and set the ReadOnly property in VolumeMounts to \"true\". If omitted, the default is \"false\". More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
                                "type": "boolean"
                              },
                              "volumeID": {
                                "description": "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
                                "type": "string"
                              }
                            }
                          },
                          "azureDisk": {
                            "description": "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
                            "type": "object",
                            "required": [
                              "diskName",
                              "diskURI"
                            ],
                            "properties": {
                              "cachingMode": {
                                "description": "Host Caching mode: None, Read Only, Read Write.",
                                "type": "string"
                              },
                              "diskName": {
                                "description": "The Name of the data disk in the blob storage",
                                "type": "string"
                              },
                              "diskURI": {
                                "description": "The URI the data disk in the blob storage",
                                "type": "string"
                              },
                              "fsType": {
                                "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
                                "type": "string"
                              },
                              "kind": {
                                "description": "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
                                "type": "string"
                              },
                              "readOnly": {
                                "description": "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
                                "type": "boolean"
                              }
                            }
                          },
                          "configMap": {
                            "description": "Adapts a ConfigMap into a volume.\n\nThe contents of the target ConfigMap's Data field will be presented in a volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. ConfigMap volumes support ownership management and SELinux relabeling.",
                            "type": "object",
                            "properties": {
                              "defaultMode": {
                                "description": "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                                "type": "integer",
                                "format": "int32"
                              },
                              "items": {
                                "description": "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
                                "type": "array",
                                "items": {
                                  "description": "Maps a string key to a path within a volume.",
                                  "type": "object",
                                  "required": [
                                    "key",
                                    "path"
                                  ],
                                  "properties": {
                                    "key": {
                                      "description": "The key to project.",
                                      "type": "string"
                                    },
                                    "mode": {
                                      "description": "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                                      "type": "integer",
                                      "format": "int32"
                                    },
                                    "path": {
                                      "description": "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
                                      "type": "string"
                                    }
                                  }
                                }
                              },
                              "name": {
                                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
                                "type": "string"
                              },
                              "optional": {
                                "description": "Specify whether the ConfigMap or it's keys must be defined",
                                "type": "boolean"
                              }
                            }
                          },
                          "emptyDir": {
                            "description": "Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.",
                            "type": "object",
                            "properties": {
                              "medium": {
                                "description": "What type of storage medium should back this directory. The default is \"\" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
                                "type": "string"
                              },
                              "sizeLimit": {
                                "description": "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/user-guide/volumes#emptydir",
                                "x-kubernetes-int-or-string": true
                              }
                            }
                          },
                          "hostPath": {
                            "description": "Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.",
                            "type": "object",
                            "required": [
                              "path"
                            ],
                            "properties": {
                              "path": {
                                "description": "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
                                "type": "string"
                              },
                              "type": {
                                "description": "Type for HostPath Volume Defaults to \"\" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
                                "type": "string"
                              }
                            }
                          },
                          "persistentVolumeClaim": {
                            "description": "PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).",
                            "type": "object",
                            "required": [
                              "claimName"
                            ],
                            "properties": {
                              "claimName": {
                                "description": "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
                                "type": "string"
                              },
                              "readOnly": {
                                "description": "Will force the ReadOnly setting in VolumeMounts. Default false.",
                                "type": "boolean"
                              }
                            }
                          }
                        }
                      },
                      "x-kubernetes-patch-merge-key": "name",
                      "x-kubernetes-patch-strategy": "merge,retainKeys"
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
`
