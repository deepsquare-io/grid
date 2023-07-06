---
title: 'Introduction'
---

# Getting Started with DeepSquare

Welcome to the Getting Started guide for DeepSquare. This tutorial aims to provide a comprehensive walkthrough to unleash the full potential of our platform. Across this guide, you will learn the essentials of configuring, deploying, and scaling High-Performance Computing (HPC) workloads using DeepSquare.

Here's an overview of the learning path we'll follow in this guide:

1. **Running a Simple 'Hello World' on DeepSquare**: Our journey begins with executing a simple 'Hello World' script on DeepSquare. This initial step will give you a hands-on experience of creating a basic workflow and understanding how DeepSquare operates.

2. **Introduction to OpenMPI**: We'll adapt the hello world example to use OpenMPI, allowing us to run the script on multiple HPC nodes! OpenMPI is an open-source implementation of the [Message Passing Interface](https://en.wikipedia.org/wiki/Message_Passing_Interface) (MPI). This powerful tool is extensively used in HPC to enable parallel computing across multiple machines or processors, providing scalability, portability, and performance optimization.

3. **Manage Storage**: Next, We'll then delve into how to handle storage with DeepSquare. It's crucial to understand that DeepSquare does not store your data; you maintain full control over your data management. However, for DeepSquare to pull and push data effectively, it requires access to your data sources.

4. **Containerizing and Running an Application**: Next, we'll delve into how to containerize an application and subsequently run it on the DeepSquare platform.

5. **Optimizing HPC Workloads for Multi-node Distributed Training**: Building on the knowledge acquired in the previous steps, we will take an example of an HPC workload - a machine learning application - and optimize it for multi-node distributed training on DeepSquare's HPC infrastructure. This process will involve using Horovod, a distributed training framework for TensorFlow, Keras, and PyTorch.

By the conclusion of this guide, you'll be equipped with the knowledge and skills required to effectively harness DeepSquare for your HPC tasks. Whether you're a researcher, data scientist, or an engineer, this guide will empower you to effectively deploy and manage HPC workloads without the burden of infrastructure maintenance.

Let's embark on this exciting journey with DeepSquare and explore how we can maximize the efficiency of your HPC tasks!
