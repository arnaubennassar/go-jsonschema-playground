# Schema Docs

- [1. Property `root > hello`](#hello)
  - [1.1. Property `root > hello > greatingMessage`](#hello_greatingMessage)
  - [1.2. Property `root > hello > greatingFrequency`](#hello_greatingFrequency)
- [2. Property `root > world`](#world)
  - [2.1. Property `root > world > greatingMessage`](#world_greatingMessage)
  - [2.2. Property `root > world > greatingFrequency`](#world_greatingFrequency)

|                           |                                                         |
| ------------------------- | ------------------------------------------------------- |
| **Type**                  | `object`                                                |
| **Required**              | No                                                      |
| **Additional properties** | [[Not allowed]](# "Additional Properties not allowed.") |

**Description:** Config is used to configure the hello world components

| Property           | Pattern | Type   | Deprecated | Definition             | Title/Description                                                                                              |
| ------------------ | ------- | ------ | ---------- | ---------------------- | -------------------------------------------------------------------------------------------------------------- |
| + [hello](#hello ) | No      | object | No         | In #/$defs/HelloConfig | Configuration of the hello component, this component will log messages on the console in a specified frequency |
| + [world](#world ) | No      | object | No         | In #/$defs/WorldConfig | Configuration of the world component, this component will log messages on the console in a specified frequency |

## <a name="hello"></a>1. Property `root > hello`

|                           |                                                         |
| ------------------------- | ------------------------------------------------------- |
| **Type**                  | `object`                                                |
| **Required**              | Yes                                                     |
| **Additional properties** | [[Not allowed]](# "Additional Properties not allowed.") |
| **Defined in**            | #/$defs/HelloConfig                                     |

**Description:** Configuration of the hello component, this component will log messages on the console in a specified frequency

| Property                                         | Pattern | Type    | Deprecated | Definition | Title/Description                                    |
| ------------------------------------------------ | ------- | ------- | ---------- | ---------- | ---------------------------------------------------- |
| + [greatingMessage](#hello_greatingMessage )     | No      | string  | No         | -          | Message that is going to be logged on the console    |
| + [greatingFrequency](#hello_greatingFrequency ) | No      | integer | No         | -          | Frequency in which the message is going to be logged |

### <a name="hello_greatingMessage"></a>1.1. Property `root > hello > greatingMessage`

|              |           |
| ------------ | --------- |
| **Type**     | `string`  |
| **Required** | Yes       |
| **Default**  | `"hello"` |

**Description:** Message that is going to be logged on the console

**Example:** 

```json
"hello world"
```

### <a name="hello_greatingFrequency"></a>1.2. Property `root > hello > greatingFrequency`

|              |           |
| ------------ | --------- |
| **Type**     | `integer` |
| **Required** | Yes       |
| **Default**  | `0`       |

**Description:** Frequency in which the message is going to be logged

## <a name="world"></a>2. Property `root > world`

|                           |                                                         |
| ------------------------- | ------------------------------------------------------- |
| **Type**                  | `object`                                                |
| **Required**              | Yes                                                     |
| **Additional properties** | [[Not allowed]](# "Additional Properties not allowed.") |
| **Defined in**            | #/$defs/WorldConfig                                     |

**Description:** Configuration of the world component, this component will log messages on the console in a specified frequency

| Property                                         | Pattern | Type    | Deprecated | Definition | Title/Description                                    |
| ------------------------------------------------ | ------- | ------- | ---------- | ---------- | ---------------------------------------------------- |
| + [greatingMessage](#world_greatingMessage )     | No      | string  | No         | -          | Message that is going to be logged on the console    |
| + [greatingFrequency](#world_greatingFrequency ) | No      | integer | No         | -          | Frequency in which the message is going to be logged |

### <a name="world_greatingMessage"></a>2.1. Property `root > world > greatingMessage`

|              |           |
| ------------ | --------- |
| **Type**     | `string`  |
| **Required** | Yes       |
| **Default**  | `"hello"` |

**Description:** Message that is going to be logged on the console

**Example:** 

```json
"hello world"
```

### <a name="world_greatingFrequency"></a>2.2. Property `root > world > greatingFrequency`

|              |           |
| ------------ | --------- |
| **Type**     | `integer` |
| **Required** | Yes       |
| **Default**  | `0`       |

**Description:** Frequency in which the message is going to be logged

----------------------------------------------------------------------------------------------------------------------------
Generated using [json-schema-for-humans](https://github.com/coveooss/json-schema-for-humans) on 2023-04-26 at 20:48:19 +0200
