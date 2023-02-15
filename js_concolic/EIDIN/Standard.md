# A Protocol for the Construction of Event-Aware DSE Engines by External Interaction

## 0. Abstract

This document defines the **EIDIN (Event-aware Interactive DSE INterface)** protocol. This is a simple two-way protocol based on Protocol Buffers. It is established between an "analyzer process" (`AN`), which generates path conditions, and an "orchestration process" (`OR`), which makes solver queries and sends models (back to the "analyzer process") to be tested. Keep in mind that this protocol is asynchronous, and the only associations between analysis requests (`Analyze*`) and responses (`PathCondition`) is an ID set by the orchestration process. 

## 1. Messages

There are six messages that make up the protocol:

| Dir.       | Message Name            | Description                                                               |
|------------|-------------------------|---------------------------------------------------------------------------|
| `OR -> AN` | `SessionInit`           | Start a new session, describes important operation parameters and target. |
| `OR -> AN` | `AnalyzeAny`            | Analyze the target program on any input (random or predefined).           |
| `OR -> AN` | `AnalyzeModel`          | Analyze the target program in accordance with a particular model.         |
| `OR <- AN` | `PathCondition`         | Provides the returned path condition (analysis results).                  |
| `OR -> AN` | `RequestCallbackSource` | Requests the source of a particular callback that analysis discovered.    |
| `OR <- AN` | `CallbackSource`        | The requested source of a particular callback that analysis discovered.   |
