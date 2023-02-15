# A Protocol for the Construction of Event-Aware DSE Engines by External Interaction

## 0. Abstract

This document defines the **EIDIN (Event-aware Interactive DSE INterface)** protocol. This is a simple two-way protocol based on Protocol Buffers. It is established between an "analyzer process", which generates path conditions, and an "orchestration process", which makes solver queries and sends models (back to the "analyzer process") to be tested. 

## 1. Messages

There are four messages that make up the protocol:

| Dir.     | Message Name    | Description                                                               |
|----------|-----------------|---------------------------------------------------------------------------|
| `A -> E` | `SessionInit`   | Start a new session, describes important operation parameters and target. |
| `A -> E` | `AnalyzeAny`    | Analyze the target program on any input (random or predefined).           |
| `A -> E` | `AnalyzeModel`  | Analyze the target program in accordance with a particular model.         |
| `A <- E` | `PathCondition` | Provides the returned path condition (analysis results).                  |
