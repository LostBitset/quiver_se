/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
(function(global, factory) { /* global define, require, module */

    /* AMD */ if (typeof define === 'function' && define.amd)
        define(["protobufjs/minimal"], factory);

    /* CommonJS */ else if (typeof require === 'function' && typeof module === 'object' && module && module.exports)
        module.exports = factory(require("protobufjs/minimal"));

})(this, function($protobuf) {
    "use strict";

    // Common aliases
    var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;
    
    // Exported root namespace
    var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});
    
    $root.eidin = (function() {
    
        /**
         * Namespace eidin.
         * @exports eidin
         * @namespace
         */
        var eidin = {};
    
        eidin.SessionInit = (function() {
    
            /**
             * Properties of a SessionInit.
             * @memberof eidin
             * @interface ISessionInit
             * @property {string|null} [eidinVersion] SessionInit eidinVersion
             * @property {string|null} [targetUri] SessionInit targetUri
             * @property {string|null} [constraintFormatOther] SessionInit constraintFormatOther
             */
    
            /**
             * Constructs a new SessionInit.
             * @memberof eidin
             * @classdesc Represents a SessionInit.
             * @implements ISessionInit
             * @constructor
             * @param {eidin.ISessionInit=} [properties] Properties to set
             */
            function SessionInit(properties) {
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }
    
            /**
             * SessionInit eidinVersion.
             * @member {string} eidinVersion
             * @memberof eidin.SessionInit
             * @instance
             */
            SessionInit.prototype.eidinVersion = "";
    
            /**
             * SessionInit targetUri.
             * @member {string} targetUri
             * @memberof eidin.SessionInit
             * @instance
             */
            SessionInit.prototype.targetUri = "";
    
            /**
             * SessionInit constraintFormatOther.
             * @member {string|null|undefined} constraintFormatOther
             * @memberof eidin.SessionInit
             * @instance
             */
            SessionInit.prototype.constraintFormatOther = null;
    
            // OneOf field names bound to virtual getters and setters
            var $oneOfFields;
    
            /**
             * SessionInit _constraintFormatOther.
             * @member {"constraintFormatOther"|undefined} _constraintFormatOther
             * @memberof eidin.SessionInit
             * @instance
             */
            Object.defineProperty(SessionInit.prototype, "_constraintFormatOther", {
                get: $util.oneOfGetter($oneOfFields = ["constraintFormatOther"]),
                set: $util.oneOfSetter($oneOfFields)
            });
    
            /**
             * Creates a new SessionInit instance using the specified properties.
             * @function create
             * @memberof eidin.SessionInit
             * @static
             * @param {eidin.ISessionInit=} [properties] Properties to set
             * @returns {eidin.SessionInit} SessionInit instance
             */
            SessionInit.create = function create(properties) {
                return new SessionInit(properties);
            };
    
            /**
             * Encodes the specified SessionInit message. Does not implicitly {@link eidin.SessionInit.verify|verify} messages.
             * @function encode
             * @memberof eidin.SessionInit
             * @static
             * @param {eidin.ISessionInit} message SessionInit message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            SessionInit.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.eidinVersion != null && Object.hasOwnProperty.call(message, "eidinVersion"))
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.eidinVersion);
                if (message.targetUri != null && Object.hasOwnProperty.call(message, "targetUri"))
                    writer.uint32(/* id 2, wireType 2 =*/18).string(message.targetUri);
                if (message.constraintFormatOther != null && Object.hasOwnProperty.call(message, "constraintFormatOther"))
                    writer.uint32(/* id 3, wireType 2 =*/26).string(message.constraintFormatOther);
                return writer;
            };
    
            /**
             * Encodes the specified SessionInit message, length delimited. Does not implicitly {@link eidin.SessionInit.verify|verify} messages.
             * @function encodeDelimited
             * @memberof eidin.SessionInit
             * @static
             * @param {eidin.ISessionInit} message SessionInit message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            SessionInit.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };
    
            /**
             * Decodes a SessionInit message from the specified reader or buffer.
             * @function decode
             * @memberof eidin.SessionInit
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {eidin.SessionInit} SessionInit
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            SessionInit.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.eidin.SessionInit();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1: {
                            message.eidinVersion = reader.string();
                            break;
                        }
                    case 2: {
                            message.targetUri = reader.string();
                            break;
                        }
                    case 3: {
                            message.constraintFormatOther = reader.string();
                            break;
                        }
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };
    
            /**
             * Decodes a SessionInit message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof eidin.SessionInit
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {eidin.SessionInit} SessionInit
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            SessionInit.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };
    
            /**
             * Verifies a SessionInit message.
             * @function verify
             * @memberof eidin.SessionInit
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            SessionInit.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                var properties = {};
                if (message.eidinVersion != null && message.hasOwnProperty("eidinVersion"))
                    if (!$util.isString(message.eidinVersion))
                        return "eidinVersion: string expected";
                if (message.targetUri != null && message.hasOwnProperty("targetUri"))
                    if (!$util.isString(message.targetUri))
                        return "targetUri: string expected";
                if (message.constraintFormatOther != null && message.hasOwnProperty("constraintFormatOther")) {
                    properties._constraintFormatOther = 1;
                    if (!$util.isString(message.constraintFormatOther))
                        return "constraintFormatOther: string expected";
                }
                return null;
            };
    
            /**
             * Creates a SessionInit message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof eidin.SessionInit
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {eidin.SessionInit} SessionInit
             */
            SessionInit.fromObject = function fromObject(object) {
                if (object instanceof $root.eidin.SessionInit)
                    return object;
                var message = new $root.eidin.SessionInit();
                if (object.eidinVersion != null)
                    message.eidinVersion = String(object.eidinVersion);
                if (object.targetUri != null)
                    message.targetUri = String(object.targetUri);
                if (object.constraintFormatOther != null)
                    message.constraintFormatOther = String(object.constraintFormatOther);
                return message;
            };
    
            /**
             * Creates a plain object from a SessionInit message. Also converts values to other types if specified.
             * @function toObject
             * @memberof eidin.SessionInit
             * @static
             * @param {eidin.SessionInit} message SessionInit
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            SessionInit.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                var object = {};
                if (options.defaults) {
                    object.eidinVersion = "";
                    object.targetUri = "";
                }
                if (message.eidinVersion != null && message.hasOwnProperty("eidinVersion"))
                    object.eidinVersion = message.eidinVersion;
                if (message.targetUri != null && message.hasOwnProperty("targetUri"))
                    object.targetUri = message.targetUri;
                if (message.constraintFormatOther != null && message.hasOwnProperty("constraintFormatOther")) {
                    object.constraintFormatOther = message.constraintFormatOther;
                    if (options.oneofs)
                        object._constraintFormatOther = "constraintFormatOther";
                }
                return object;
            };
    
            /**
             * Converts this SessionInit to JSON.
             * @function toJSON
             * @memberof eidin.SessionInit
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            SessionInit.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };
    
            /**
             * Gets the default type url for SessionInit
             * @function getTypeUrl
             * @memberof eidin.SessionInit
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            SessionInit.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/eidin.SessionInit";
            };
    
            /**
             * ConstraintFormat enum.
             * @name eidin.SessionInit.ConstraintFormat
             * @enum {number}
             * @property {number} SMTLIB_V2=0 SMTLIB_V2 value
             * @property {number} SMTLIB_2VA=1 SMTLIB_2VA value
             * @property {number} OTHER=2 OTHER value
             */
            SessionInit.ConstraintFormat = (function() {
                var valuesById = {}, values = Object.create(valuesById);
                values[valuesById[0] = "SMTLIB_V2"] = 0;
                values[valuesById[1] = "SMTLIB_2VA"] = 1;
                values[valuesById[2] = "OTHER"] = 2;
                return values;
            })();
    
            return SessionInit;
        })();
    
        eidin.AnalyzeAny = (function() {
    
            /**
             * Properties of an AnalyzeAny.
             * @memberof eidin
             * @interface IAnalyzeAny
             * @property {boolean|null} [forbidCaching] AnalyzeAny forbidCaching
             */
    
            /**
             * Constructs a new AnalyzeAny.
             * @memberof eidin
             * @classdesc Represents an AnalyzeAny.
             * @implements IAnalyzeAny
             * @constructor
             * @param {eidin.IAnalyzeAny=} [properties] Properties to set
             */
            function AnalyzeAny(properties) {
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }
    
            /**
             * AnalyzeAny forbidCaching.
             * @member {boolean} forbidCaching
             * @memberof eidin.AnalyzeAny
             * @instance
             */
            AnalyzeAny.prototype.forbidCaching = false;
    
            /**
             * Creates a new AnalyzeAny instance using the specified properties.
             * @function create
             * @memberof eidin.AnalyzeAny
             * @static
             * @param {eidin.IAnalyzeAny=} [properties] Properties to set
             * @returns {eidin.AnalyzeAny} AnalyzeAny instance
             */
            AnalyzeAny.create = function create(properties) {
                return new AnalyzeAny(properties);
            };
    
            /**
             * Encodes the specified AnalyzeAny message. Does not implicitly {@link eidin.AnalyzeAny.verify|verify} messages.
             * @function encode
             * @memberof eidin.AnalyzeAny
             * @static
             * @param {eidin.IAnalyzeAny} message AnalyzeAny message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            AnalyzeAny.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.forbidCaching != null && Object.hasOwnProperty.call(message, "forbidCaching"))
                    writer.uint32(/* id 1, wireType 0 =*/8).bool(message.forbidCaching);
                return writer;
            };
    
            /**
             * Encodes the specified AnalyzeAny message, length delimited. Does not implicitly {@link eidin.AnalyzeAny.verify|verify} messages.
             * @function encodeDelimited
             * @memberof eidin.AnalyzeAny
             * @static
             * @param {eidin.IAnalyzeAny} message AnalyzeAny message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            AnalyzeAny.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };
    
            /**
             * Decodes an AnalyzeAny message from the specified reader or buffer.
             * @function decode
             * @memberof eidin.AnalyzeAny
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {eidin.AnalyzeAny} AnalyzeAny
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            AnalyzeAny.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.eidin.AnalyzeAny();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1: {
                            message.forbidCaching = reader.bool();
                            break;
                        }
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };
    
            /**
             * Decodes an AnalyzeAny message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof eidin.AnalyzeAny
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {eidin.AnalyzeAny} AnalyzeAny
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            AnalyzeAny.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };
    
            /**
             * Verifies an AnalyzeAny message.
             * @function verify
             * @memberof eidin.AnalyzeAny
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            AnalyzeAny.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.forbidCaching != null && message.hasOwnProperty("forbidCaching"))
                    if (typeof message.forbidCaching !== "boolean")
                        return "forbidCaching: boolean expected";
                return null;
            };
    
            /**
             * Creates an AnalyzeAny message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof eidin.AnalyzeAny
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {eidin.AnalyzeAny} AnalyzeAny
             */
            AnalyzeAny.fromObject = function fromObject(object) {
                if (object instanceof $root.eidin.AnalyzeAny)
                    return object;
                var message = new $root.eidin.AnalyzeAny();
                if (object.forbidCaching != null)
                    message.forbidCaching = Boolean(object.forbidCaching);
                return message;
            };
    
            /**
             * Creates a plain object from an AnalyzeAny message. Also converts values to other types if specified.
             * @function toObject
             * @memberof eidin.AnalyzeAny
             * @static
             * @param {eidin.AnalyzeAny} message AnalyzeAny
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            AnalyzeAny.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                var object = {};
                if (options.defaults)
                    object.forbidCaching = false;
                if (message.forbidCaching != null && message.hasOwnProperty("forbidCaching"))
                    object.forbidCaching = message.forbidCaching;
                return object;
            };
    
            /**
             * Converts this AnalyzeAny to JSON.
             * @function toJSON
             * @memberof eidin.AnalyzeAny
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            AnalyzeAny.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };
    
            /**
             * Gets the default type url for AnalyzeAny
             * @function getTypeUrl
             * @memberof eidin.AnalyzeAny
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            AnalyzeAny.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/eidin.AnalyzeAny";
            };
    
            return AnalyzeAny;
        })();
    
        eidin.AnalyzeModel = (function() {
    
            /**
             * Properties of an AnalyzeModel.
             * @memberof eidin
             * @interface IAnalyzeModel
             * @property {boolean|null} [forbidCaching] AnalyzeModel forbidCaching
             * @property {string|null} [model] AnalyzeModel model
             */
    
            /**
             * Constructs a new AnalyzeModel.
             * @memberof eidin
             * @classdesc Represents an AnalyzeModel.
             * @implements IAnalyzeModel
             * @constructor
             * @param {eidin.IAnalyzeModel=} [properties] Properties to set
             */
            function AnalyzeModel(properties) {
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }
    
            /**
             * AnalyzeModel forbidCaching.
             * @member {boolean} forbidCaching
             * @memberof eidin.AnalyzeModel
             * @instance
             */
            AnalyzeModel.prototype.forbidCaching = false;
    
            /**
             * AnalyzeModel model.
             * @member {string} model
             * @memberof eidin.AnalyzeModel
             * @instance
             */
            AnalyzeModel.prototype.model = "";
    
            /**
             * Creates a new AnalyzeModel instance using the specified properties.
             * @function create
             * @memberof eidin.AnalyzeModel
             * @static
             * @param {eidin.IAnalyzeModel=} [properties] Properties to set
             * @returns {eidin.AnalyzeModel} AnalyzeModel instance
             */
            AnalyzeModel.create = function create(properties) {
                return new AnalyzeModel(properties);
            };
    
            /**
             * Encodes the specified AnalyzeModel message. Does not implicitly {@link eidin.AnalyzeModel.verify|verify} messages.
             * @function encode
             * @memberof eidin.AnalyzeModel
             * @static
             * @param {eidin.IAnalyzeModel} message AnalyzeModel message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            AnalyzeModel.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.forbidCaching != null && Object.hasOwnProperty.call(message, "forbidCaching"))
                    writer.uint32(/* id 1, wireType 0 =*/8).bool(message.forbidCaching);
                if (message.model != null && Object.hasOwnProperty.call(message, "model"))
                    writer.uint32(/* id 2, wireType 2 =*/18).string(message.model);
                return writer;
            };
    
            /**
             * Encodes the specified AnalyzeModel message, length delimited. Does not implicitly {@link eidin.AnalyzeModel.verify|verify} messages.
             * @function encodeDelimited
             * @memberof eidin.AnalyzeModel
             * @static
             * @param {eidin.IAnalyzeModel} message AnalyzeModel message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            AnalyzeModel.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };
    
            /**
             * Decodes an AnalyzeModel message from the specified reader or buffer.
             * @function decode
             * @memberof eidin.AnalyzeModel
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {eidin.AnalyzeModel} AnalyzeModel
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            AnalyzeModel.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.eidin.AnalyzeModel();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1: {
                            message.forbidCaching = reader.bool();
                            break;
                        }
                    case 2: {
                            message.model = reader.string();
                            break;
                        }
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };
    
            /**
             * Decodes an AnalyzeModel message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof eidin.AnalyzeModel
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {eidin.AnalyzeModel} AnalyzeModel
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            AnalyzeModel.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };
    
            /**
             * Verifies an AnalyzeModel message.
             * @function verify
             * @memberof eidin.AnalyzeModel
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            AnalyzeModel.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.forbidCaching != null && message.hasOwnProperty("forbidCaching"))
                    if (typeof message.forbidCaching !== "boolean")
                        return "forbidCaching: boolean expected";
                if (message.model != null && message.hasOwnProperty("model"))
                    if (!$util.isString(message.model))
                        return "model: string expected";
                return null;
            };
    
            /**
             * Creates an AnalyzeModel message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof eidin.AnalyzeModel
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {eidin.AnalyzeModel} AnalyzeModel
             */
            AnalyzeModel.fromObject = function fromObject(object) {
                if (object instanceof $root.eidin.AnalyzeModel)
                    return object;
                var message = new $root.eidin.AnalyzeModel();
                if (object.forbidCaching != null)
                    message.forbidCaching = Boolean(object.forbidCaching);
                if (object.model != null)
                    message.model = String(object.model);
                return message;
            };
    
            /**
             * Creates a plain object from an AnalyzeModel message. Also converts values to other types if specified.
             * @function toObject
             * @memberof eidin.AnalyzeModel
             * @static
             * @param {eidin.AnalyzeModel} message AnalyzeModel
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            AnalyzeModel.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                var object = {};
                if (options.defaults) {
                    object.forbidCaching = false;
                    object.model = "";
                }
                if (message.forbidCaching != null && message.hasOwnProperty("forbidCaching"))
                    object.forbidCaching = message.forbidCaching;
                if (message.model != null && message.hasOwnProperty("model"))
                    object.model = message.model;
                return object;
            };
    
            /**
             * Converts this AnalyzeModel to JSON.
             * @function toJSON
             * @memberof eidin.AnalyzeModel
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            AnalyzeModel.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };
    
            /**
             * Gets the default type url for AnalyzeModel
             * @function getTypeUrl
             * @memberof eidin.AnalyzeModel
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            AnalyzeModel.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/eidin.AnalyzeModel";
            };
    
            return AnalyzeModel;
        })();
    
        eidin.PathCondition = (function() {
    
            /**
             * Properties of a PathCondition.
             * @memberof eidin
             * @interface IPathCondition
             * @property {Array.<eidin.ISMTFreeFun>|null} [freeFuns] PathCondition freeFuns
             * @property {Array.<eidin.IPathConditionSegment>|null} [segmentedPc] PathCondition segmentedPc
             */
    
            /**
             * Constructs a new PathCondition.
             * @memberof eidin
             * @classdesc Represents a PathCondition.
             * @implements IPathCondition
             * @constructor
             * @param {eidin.IPathCondition=} [properties] Properties to set
             */
            function PathCondition(properties) {
                this.freeFuns = [];
                this.segmentedPc = [];
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }
    
            /**
             * PathCondition freeFuns.
             * @member {Array.<eidin.ISMTFreeFun>} freeFuns
             * @memberof eidin.PathCondition
             * @instance
             */
            PathCondition.prototype.freeFuns = $util.emptyArray;
    
            /**
             * PathCondition segmentedPc.
             * @member {Array.<eidin.IPathConditionSegment>} segmentedPc
             * @memberof eidin.PathCondition
             * @instance
             */
            PathCondition.prototype.segmentedPc = $util.emptyArray;
    
            /**
             * Creates a new PathCondition instance using the specified properties.
             * @function create
             * @memberof eidin.PathCondition
             * @static
             * @param {eidin.IPathCondition=} [properties] Properties to set
             * @returns {eidin.PathCondition} PathCondition instance
             */
            PathCondition.create = function create(properties) {
                return new PathCondition(properties);
            };
    
            /**
             * Encodes the specified PathCondition message. Does not implicitly {@link eidin.PathCondition.verify|verify} messages.
             * @function encode
             * @memberof eidin.PathCondition
             * @static
             * @param {eidin.IPathCondition} message PathCondition message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            PathCondition.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.freeFuns != null && message.freeFuns.length)
                    for (var i = 0; i < message.freeFuns.length; ++i)
                        $root.eidin.SMTFreeFun.encode(message.freeFuns[i], writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
                if (message.segmentedPc != null && message.segmentedPc.length)
                    for (var i = 0; i < message.segmentedPc.length; ++i)
                        $root.eidin.PathConditionSegment.encode(message.segmentedPc[i], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                return writer;
            };
    
            /**
             * Encodes the specified PathCondition message, length delimited. Does not implicitly {@link eidin.PathCondition.verify|verify} messages.
             * @function encodeDelimited
             * @memberof eidin.PathCondition
             * @static
             * @param {eidin.IPathCondition} message PathCondition message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            PathCondition.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };
    
            /**
             * Decodes a PathCondition message from the specified reader or buffer.
             * @function decode
             * @memberof eidin.PathCondition
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {eidin.PathCondition} PathCondition
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            PathCondition.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.eidin.PathCondition();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1: {
                            if (!(message.freeFuns && message.freeFuns.length))
                                message.freeFuns = [];
                            message.freeFuns.push($root.eidin.SMTFreeFun.decode(reader, reader.uint32()));
                            break;
                        }
                    case 2: {
                            if (!(message.segmentedPc && message.segmentedPc.length))
                                message.segmentedPc = [];
                            message.segmentedPc.push($root.eidin.PathConditionSegment.decode(reader, reader.uint32()));
                            break;
                        }
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };
    
            /**
             * Decodes a PathCondition message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof eidin.PathCondition
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {eidin.PathCondition} PathCondition
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            PathCondition.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };
    
            /**
             * Verifies a PathCondition message.
             * @function verify
             * @memberof eidin.PathCondition
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            PathCondition.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.freeFuns != null && message.hasOwnProperty("freeFuns")) {
                    if (!Array.isArray(message.freeFuns))
                        return "freeFuns: array expected";
                    for (var i = 0; i < message.freeFuns.length; ++i) {
                        var error = $root.eidin.SMTFreeFun.verify(message.freeFuns[i]);
                        if (error)
                            return "freeFuns." + error;
                    }
                }
                if (message.segmentedPc != null && message.hasOwnProperty("segmentedPc")) {
                    if (!Array.isArray(message.segmentedPc))
                        return "segmentedPc: array expected";
                    for (var i = 0; i < message.segmentedPc.length; ++i) {
                        var error = $root.eidin.PathConditionSegment.verify(message.segmentedPc[i]);
                        if (error)
                            return "segmentedPc." + error;
                    }
                }
                return null;
            };
    
            /**
             * Creates a PathCondition message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof eidin.PathCondition
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {eidin.PathCondition} PathCondition
             */
            PathCondition.fromObject = function fromObject(object) {
                if (object instanceof $root.eidin.PathCondition)
                    return object;
                var message = new $root.eidin.PathCondition();
                if (object.freeFuns) {
                    if (!Array.isArray(object.freeFuns))
                        throw TypeError(".eidin.PathCondition.freeFuns: array expected");
                    message.freeFuns = [];
                    for (var i = 0; i < object.freeFuns.length; ++i) {
                        if (typeof object.freeFuns[i] !== "object")
                            throw TypeError(".eidin.PathCondition.freeFuns: object expected");
                        message.freeFuns[i] = $root.eidin.SMTFreeFun.fromObject(object.freeFuns[i]);
                    }
                }
                if (object.segmentedPc) {
                    if (!Array.isArray(object.segmentedPc))
                        throw TypeError(".eidin.PathCondition.segmentedPc: array expected");
                    message.segmentedPc = [];
                    for (var i = 0; i < object.segmentedPc.length; ++i) {
                        if (typeof object.segmentedPc[i] !== "object")
                            throw TypeError(".eidin.PathCondition.segmentedPc: object expected");
                        message.segmentedPc[i] = $root.eidin.PathConditionSegment.fromObject(object.segmentedPc[i]);
                    }
                }
                return message;
            };
    
            /**
             * Creates a plain object from a PathCondition message. Also converts values to other types if specified.
             * @function toObject
             * @memberof eidin.PathCondition
             * @static
             * @param {eidin.PathCondition} message PathCondition
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            PathCondition.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                var object = {};
                if (options.arrays || options.defaults) {
                    object.freeFuns = [];
                    object.segmentedPc = [];
                }
                if (message.freeFuns && message.freeFuns.length) {
                    object.freeFuns = [];
                    for (var j = 0; j < message.freeFuns.length; ++j)
                        object.freeFuns[j] = $root.eidin.SMTFreeFun.toObject(message.freeFuns[j], options);
                }
                if (message.segmentedPc && message.segmentedPc.length) {
                    object.segmentedPc = [];
                    for (var j = 0; j < message.segmentedPc.length; ++j)
                        object.segmentedPc[j] = $root.eidin.PathConditionSegment.toObject(message.segmentedPc[j], options);
                }
                return object;
            };
    
            /**
             * Converts this PathCondition to JSON.
             * @function toJSON
             * @memberof eidin.PathCondition
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            PathCondition.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };
    
            /**
             * Gets the default type url for PathCondition
             * @function getTypeUrl
             * @memberof eidin.PathCondition
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            PathCondition.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/eidin.PathCondition";
            };
    
            return PathCondition;
        })();
    
        eidin.SMTFreeFun = (function() {
    
            /**
             * Properties of a SMTFreeFun.
             * @memberof eidin
             * @interface ISMTFreeFun
             * @property {string|null} [name] SMTFreeFun name
             * @property {Array.<string>|null} [argSorts] SMTFreeFun argSorts
             * @property {string|null} [retSort] SMTFreeFun retSort
             */
    
            /**
             * Constructs a new SMTFreeFun.
             * @memberof eidin
             * @classdesc Represents a SMTFreeFun.
             * @implements ISMTFreeFun
             * @constructor
             * @param {eidin.ISMTFreeFun=} [properties] Properties to set
             */
            function SMTFreeFun(properties) {
                this.argSorts = [];
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }
    
            /**
             * SMTFreeFun name.
             * @member {string} name
             * @memberof eidin.SMTFreeFun
             * @instance
             */
            SMTFreeFun.prototype.name = "";
    
            /**
             * SMTFreeFun argSorts.
             * @member {Array.<string>} argSorts
             * @memberof eidin.SMTFreeFun
             * @instance
             */
            SMTFreeFun.prototype.argSorts = $util.emptyArray;
    
            /**
             * SMTFreeFun retSort.
             * @member {string} retSort
             * @memberof eidin.SMTFreeFun
             * @instance
             */
            SMTFreeFun.prototype.retSort = "";
    
            /**
             * Creates a new SMTFreeFun instance using the specified properties.
             * @function create
             * @memberof eidin.SMTFreeFun
             * @static
             * @param {eidin.ISMTFreeFun=} [properties] Properties to set
             * @returns {eidin.SMTFreeFun} SMTFreeFun instance
             */
            SMTFreeFun.create = function create(properties) {
                return new SMTFreeFun(properties);
            };
    
            /**
             * Encodes the specified SMTFreeFun message. Does not implicitly {@link eidin.SMTFreeFun.verify|verify} messages.
             * @function encode
             * @memberof eidin.SMTFreeFun
             * @static
             * @param {eidin.ISMTFreeFun} message SMTFreeFun message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            SMTFreeFun.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.name != null && Object.hasOwnProperty.call(message, "name"))
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.name);
                if (message.argSorts != null && message.argSorts.length)
                    for (var i = 0; i < message.argSorts.length; ++i)
                        writer.uint32(/* id 2, wireType 2 =*/18).string(message.argSorts[i]);
                if (message.retSort != null && Object.hasOwnProperty.call(message, "retSort"))
                    writer.uint32(/* id 3, wireType 2 =*/26).string(message.retSort);
                return writer;
            };
    
            /**
             * Encodes the specified SMTFreeFun message, length delimited. Does not implicitly {@link eidin.SMTFreeFun.verify|verify} messages.
             * @function encodeDelimited
             * @memberof eidin.SMTFreeFun
             * @static
             * @param {eidin.ISMTFreeFun} message SMTFreeFun message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            SMTFreeFun.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };
    
            /**
             * Decodes a SMTFreeFun message from the specified reader or buffer.
             * @function decode
             * @memberof eidin.SMTFreeFun
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {eidin.SMTFreeFun} SMTFreeFun
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            SMTFreeFun.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.eidin.SMTFreeFun();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1: {
                            message.name = reader.string();
                            break;
                        }
                    case 2: {
                            if (!(message.argSorts && message.argSorts.length))
                                message.argSorts = [];
                            message.argSorts.push(reader.string());
                            break;
                        }
                    case 3: {
                            message.retSort = reader.string();
                            break;
                        }
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };
    
            /**
             * Decodes a SMTFreeFun message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof eidin.SMTFreeFun
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {eidin.SMTFreeFun} SMTFreeFun
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            SMTFreeFun.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };
    
            /**
             * Verifies a SMTFreeFun message.
             * @function verify
             * @memberof eidin.SMTFreeFun
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            SMTFreeFun.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.name != null && message.hasOwnProperty("name"))
                    if (!$util.isString(message.name))
                        return "name: string expected";
                if (message.argSorts != null && message.hasOwnProperty("argSorts")) {
                    if (!Array.isArray(message.argSorts))
                        return "argSorts: array expected";
                    for (var i = 0; i < message.argSorts.length; ++i)
                        if (!$util.isString(message.argSorts[i]))
                            return "argSorts: string[] expected";
                }
                if (message.retSort != null && message.hasOwnProperty("retSort"))
                    if (!$util.isString(message.retSort))
                        return "retSort: string expected";
                return null;
            };
    
            /**
             * Creates a SMTFreeFun message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof eidin.SMTFreeFun
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {eidin.SMTFreeFun} SMTFreeFun
             */
            SMTFreeFun.fromObject = function fromObject(object) {
                if (object instanceof $root.eidin.SMTFreeFun)
                    return object;
                var message = new $root.eidin.SMTFreeFun();
                if (object.name != null)
                    message.name = String(object.name);
                if (object.argSorts) {
                    if (!Array.isArray(object.argSorts))
                        throw TypeError(".eidin.SMTFreeFun.argSorts: array expected");
                    message.argSorts = [];
                    for (var i = 0; i < object.argSorts.length; ++i)
                        message.argSorts[i] = String(object.argSorts[i]);
                }
                if (object.retSort != null)
                    message.retSort = String(object.retSort);
                return message;
            };
    
            /**
             * Creates a plain object from a SMTFreeFun message. Also converts values to other types if specified.
             * @function toObject
             * @memberof eidin.SMTFreeFun
             * @static
             * @param {eidin.SMTFreeFun} message SMTFreeFun
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            SMTFreeFun.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                var object = {};
                if (options.arrays || options.defaults)
                    object.argSorts = [];
                if (options.defaults) {
                    object.name = "";
                    object.retSort = "";
                }
                if (message.name != null && message.hasOwnProperty("name"))
                    object.name = message.name;
                if (message.argSorts && message.argSorts.length) {
                    object.argSorts = [];
                    for (var j = 0; j < message.argSorts.length; ++j)
                        object.argSorts[j] = message.argSorts[j];
                }
                if (message.retSort != null && message.hasOwnProperty("retSort"))
                    object.retSort = message.retSort;
                return object;
            };
    
            /**
             * Converts this SMTFreeFun to JSON.
             * @function toJSON
             * @memberof eidin.SMTFreeFun
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            SMTFreeFun.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };
    
            /**
             * Gets the default type url for SMTFreeFun
             * @function getTypeUrl
             * @memberof eidin.SMTFreeFun
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            SMTFreeFun.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/eidin.SMTFreeFun";
            };
    
            return SMTFreeFun;
        })();
    
        eidin.PathConditionSegment = (function() {
    
            /**
             * Properties of a PathConditionSegment.
             * @memberof eidin
             * @interface IPathConditionSegment
             * @property {eidin.ICallbackId|null} [thisCallbackId] PathConditionSegment thisCallbackId
             * @property {eidin.ICallbackId|null} [nextCallbackId] PathConditionSegment nextCallbackId
             * @property {Array.<eidin.ISMTConstraint>|null} [partialPc] PathConditionSegment partialPc
             */
    
            /**
             * Constructs a new PathConditionSegment.
             * @memberof eidin
             * @classdesc Represents a PathConditionSegment.
             * @implements IPathConditionSegment
             * @constructor
             * @param {eidin.IPathConditionSegment=} [properties] Properties to set
             */
            function PathConditionSegment(properties) {
                this.partialPc = [];
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }
    
            /**
             * PathConditionSegment thisCallbackId.
             * @member {eidin.ICallbackId|null|undefined} thisCallbackId
             * @memberof eidin.PathConditionSegment
             * @instance
             */
            PathConditionSegment.prototype.thisCallbackId = null;
    
            /**
             * PathConditionSegment nextCallbackId.
             * @member {eidin.ICallbackId|null|undefined} nextCallbackId
             * @memberof eidin.PathConditionSegment
             * @instance
             */
            PathConditionSegment.prototype.nextCallbackId = null;
    
            /**
             * PathConditionSegment partialPc.
             * @member {Array.<eidin.ISMTConstraint>} partialPc
             * @memberof eidin.PathConditionSegment
             * @instance
             */
            PathConditionSegment.prototype.partialPc = $util.emptyArray;
    
            // OneOf field names bound to virtual getters and setters
            var $oneOfFields;
    
            /**
             * PathConditionSegment _nextCallbackId.
             * @member {"nextCallbackId"|undefined} _nextCallbackId
             * @memberof eidin.PathConditionSegment
             * @instance
             */
            Object.defineProperty(PathConditionSegment.prototype, "_nextCallbackId", {
                get: $util.oneOfGetter($oneOfFields = ["nextCallbackId"]),
                set: $util.oneOfSetter($oneOfFields)
            });
    
            /**
             * Creates a new PathConditionSegment instance using the specified properties.
             * @function create
             * @memberof eidin.PathConditionSegment
             * @static
             * @param {eidin.IPathConditionSegment=} [properties] Properties to set
             * @returns {eidin.PathConditionSegment} PathConditionSegment instance
             */
            PathConditionSegment.create = function create(properties) {
                return new PathConditionSegment(properties);
            };
    
            /**
             * Encodes the specified PathConditionSegment message. Does not implicitly {@link eidin.PathConditionSegment.verify|verify} messages.
             * @function encode
             * @memberof eidin.PathConditionSegment
             * @static
             * @param {eidin.IPathConditionSegment} message PathConditionSegment message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            PathConditionSegment.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.thisCallbackId != null && Object.hasOwnProperty.call(message, "thisCallbackId"))
                    $root.eidin.CallbackId.encode(message.thisCallbackId, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
                if (message.nextCallbackId != null && Object.hasOwnProperty.call(message, "nextCallbackId"))
                    $root.eidin.CallbackId.encode(message.nextCallbackId, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                if (message.partialPc != null && message.partialPc.length)
                    for (var i = 0; i < message.partialPc.length; ++i)
                        $root.eidin.SMTConstraint.encode(message.partialPc[i], writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
                return writer;
            };
    
            /**
             * Encodes the specified PathConditionSegment message, length delimited. Does not implicitly {@link eidin.PathConditionSegment.verify|verify} messages.
             * @function encodeDelimited
             * @memberof eidin.PathConditionSegment
             * @static
             * @param {eidin.IPathConditionSegment} message PathConditionSegment message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            PathConditionSegment.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };
    
            /**
             * Decodes a PathConditionSegment message from the specified reader or buffer.
             * @function decode
             * @memberof eidin.PathConditionSegment
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {eidin.PathConditionSegment} PathConditionSegment
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            PathConditionSegment.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.eidin.PathConditionSegment();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1: {
                            message.thisCallbackId = $root.eidin.CallbackId.decode(reader, reader.uint32());
                            break;
                        }
                    case 2: {
                            message.nextCallbackId = $root.eidin.CallbackId.decode(reader, reader.uint32());
                            break;
                        }
                    case 3: {
                            if (!(message.partialPc && message.partialPc.length))
                                message.partialPc = [];
                            message.partialPc.push($root.eidin.SMTConstraint.decode(reader, reader.uint32()));
                            break;
                        }
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };
    
            /**
             * Decodes a PathConditionSegment message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof eidin.PathConditionSegment
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {eidin.PathConditionSegment} PathConditionSegment
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            PathConditionSegment.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };
    
            /**
             * Verifies a PathConditionSegment message.
             * @function verify
             * @memberof eidin.PathConditionSegment
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            PathConditionSegment.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                var properties = {};
                if (message.thisCallbackId != null && message.hasOwnProperty("thisCallbackId")) {
                    var error = $root.eidin.CallbackId.verify(message.thisCallbackId);
                    if (error)
                        return "thisCallbackId." + error;
                }
                if (message.nextCallbackId != null && message.hasOwnProperty("nextCallbackId")) {
                    properties._nextCallbackId = 1;
                    {
                        var error = $root.eidin.CallbackId.verify(message.nextCallbackId);
                        if (error)
                            return "nextCallbackId." + error;
                    }
                }
                if (message.partialPc != null && message.hasOwnProperty("partialPc")) {
                    if (!Array.isArray(message.partialPc))
                        return "partialPc: array expected";
                    for (var i = 0; i < message.partialPc.length; ++i) {
                        var error = $root.eidin.SMTConstraint.verify(message.partialPc[i]);
                        if (error)
                            return "partialPc." + error;
                    }
                }
                return null;
            };
    
            /**
             * Creates a PathConditionSegment message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof eidin.PathConditionSegment
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {eidin.PathConditionSegment} PathConditionSegment
             */
            PathConditionSegment.fromObject = function fromObject(object) {
                if (object instanceof $root.eidin.PathConditionSegment)
                    return object;
                var message = new $root.eidin.PathConditionSegment();
                if (object.thisCallbackId != null) {
                    if (typeof object.thisCallbackId !== "object")
                        throw TypeError(".eidin.PathConditionSegment.thisCallbackId: object expected");
                    message.thisCallbackId = $root.eidin.CallbackId.fromObject(object.thisCallbackId);
                }
                if (object.nextCallbackId != null) {
                    if (typeof object.nextCallbackId !== "object")
                        throw TypeError(".eidin.PathConditionSegment.nextCallbackId: object expected");
                    message.nextCallbackId = $root.eidin.CallbackId.fromObject(object.nextCallbackId);
                }
                if (object.partialPc) {
                    if (!Array.isArray(object.partialPc))
                        throw TypeError(".eidin.PathConditionSegment.partialPc: array expected");
                    message.partialPc = [];
                    for (var i = 0; i < object.partialPc.length; ++i) {
                        if (typeof object.partialPc[i] !== "object")
                            throw TypeError(".eidin.PathConditionSegment.partialPc: object expected");
                        message.partialPc[i] = $root.eidin.SMTConstraint.fromObject(object.partialPc[i]);
                    }
                }
                return message;
            };
    
            /**
             * Creates a plain object from a PathConditionSegment message. Also converts values to other types if specified.
             * @function toObject
             * @memberof eidin.PathConditionSegment
             * @static
             * @param {eidin.PathConditionSegment} message PathConditionSegment
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            PathConditionSegment.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                var object = {};
                if (options.arrays || options.defaults)
                    object.partialPc = [];
                if (options.defaults)
                    object.thisCallbackId = null;
                if (message.thisCallbackId != null && message.hasOwnProperty("thisCallbackId"))
                    object.thisCallbackId = $root.eidin.CallbackId.toObject(message.thisCallbackId, options);
                if (message.nextCallbackId != null && message.hasOwnProperty("nextCallbackId")) {
                    object.nextCallbackId = $root.eidin.CallbackId.toObject(message.nextCallbackId, options);
                    if (options.oneofs)
                        object._nextCallbackId = "nextCallbackId";
                }
                if (message.partialPc && message.partialPc.length) {
                    object.partialPc = [];
                    for (var j = 0; j < message.partialPc.length; ++j)
                        object.partialPc[j] = $root.eidin.SMTConstraint.toObject(message.partialPc[j], options);
                }
                return object;
            };
    
            /**
             * Converts this PathConditionSegment to JSON.
             * @function toJSON
             * @memberof eidin.PathConditionSegment
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            PathConditionSegment.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };
    
            /**
             * Gets the default type url for PathConditionSegment
             * @function getTypeUrl
             * @memberof eidin.PathConditionSegment
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            PathConditionSegment.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/eidin.PathConditionSegment";
            };
    
            return PathConditionSegment;
        })();
    
        eidin.SMTConstraint = (function() {
    
            /**
             * Properties of a SMTConstraint.
             * @memberof eidin
             * @interface ISMTConstraint
             * @property {string|null} [constraint] SMTConstraint constraint
             * @property {boolean|null} [assertionValue] SMTConstraint assertionValue
             */
    
            /**
             * Constructs a new SMTConstraint.
             * @memberof eidin
             * @classdesc Represents a SMTConstraint.
             * @implements ISMTConstraint
             * @constructor
             * @param {eidin.ISMTConstraint=} [properties] Properties to set
             */
            function SMTConstraint(properties) {
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }
    
            /**
             * SMTConstraint constraint.
             * @member {string} constraint
             * @memberof eidin.SMTConstraint
             * @instance
             */
            SMTConstraint.prototype.constraint = "";
    
            /**
             * SMTConstraint assertionValue.
             * @member {boolean|null|undefined} assertionValue
             * @memberof eidin.SMTConstraint
             * @instance
             */
            SMTConstraint.prototype.assertionValue = null;
    
            // OneOf field names bound to virtual getters and setters
            var $oneOfFields;
    
            /**
             * SMTConstraint _assertionValue.
             * @member {"assertionValue"|undefined} _assertionValue
             * @memberof eidin.SMTConstraint
             * @instance
             */
            Object.defineProperty(SMTConstraint.prototype, "_assertionValue", {
                get: $util.oneOfGetter($oneOfFields = ["assertionValue"]),
                set: $util.oneOfSetter($oneOfFields)
            });
    
            /**
             * Creates a new SMTConstraint instance using the specified properties.
             * @function create
             * @memberof eidin.SMTConstraint
             * @static
             * @param {eidin.ISMTConstraint=} [properties] Properties to set
             * @returns {eidin.SMTConstraint} SMTConstraint instance
             */
            SMTConstraint.create = function create(properties) {
                return new SMTConstraint(properties);
            };
    
            /**
             * Encodes the specified SMTConstraint message. Does not implicitly {@link eidin.SMTConstraint.verify|verify} messages.
             * @function encode
             * @memberof eidin.SMTConstraint
             * @static
             * @param {eidin.ISMTConstraint} message SMTConstraint message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            SMTConstraint.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.constraint != null && Object.hasOwnProperty.call(message, "constraint"))
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.constraint);
                if (message.assertionValue != null && Object.hasOwnProperty.call(message, "assertionValue"))
                    writer.uint32(/* id 2, wireType 0 =*/16).bool(message.assertionValue);
                return writer;
            };
    
            /**
             * Encodes the specified SMTConstraint message, length delimited. Does not implicitly {@link eidin.SMTConstraint.verify|verify} messages.
             * @function encodeDelimited
             * @memberof eidin.SMTConstraint
             * @static
             * @param {eidin.ISMTConstraint} message SMTConstraint message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            SMTConstraint.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };
    
            /**
             * Decodes a SMTConstraint message from the specified reader or buffer.
             * @function decode
             * @memberof eidin.SMTConstraint
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {eidin.SMTConstraint} SMTConstraint
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            SMTConstraint.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.eidin.SMTConstraint();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1: {
                            message.constraint = reader.string();
                            break;
                        }
                    case 2: {
                            message.assertionValue = reader.bool();
                            break;
                        }
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };
    
            /**
             * Decodes a SMTConstraint message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof eidin.SMTConstraint
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {eidin.SMTConstraint} SMTConstraint
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            SMTConstraint.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };
    
            /**
             * Verifies a SMTConstraint message.
             * @function verify
             * @memberof eidin.SMTConstraint
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            SMTConstraint.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                var properties = {};
                if (message.constraint != null && message.hasOwnProperty("constraint"))
                    if (!$util.isString(message.constraint))
                        return "constraint: string expected";
                if (message.assertionValue != null && message.hasOwnProperty("assertionValue")) {
                    properties._assertionValue = 1;
                    if (typeof message.assertionValue !== "boolean")
                        return "assertionValue: boolean expected";
                }
                return null;
            };
    
            /**
             * Creates a SMTConstraint message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof eidin.SMTConstraint
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {eidin.SMTConstraint} SMTConstraint
             */
            SMTConstraint.fromObject = function fromObject(object) {
                if (object instanceof $root.eidin.SMTConstraint)
                    return object;
                var message = new $root.eidin.SMTConstraint();
                if (object.constraint != null)
                    message.constraint = String(object.constraint);
                if (object.assertionValue != null)
                    message.assertionValue = Boolean(object.assertionValue);
                return message;
            };
    
            /**
             * Creates a plain object from a SMTConstraint message. Also converts values to other types if specified.
             * @function toObject
             * @memberof eidin.SMTConstraint
             * @static
             * @param {eidin.SMTConstraint} message SMTConstraint
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            SMTConstraint.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                var object = {};
                if (options.defaults)
                    object.constraint = "";
                if (message.constraint != null && message.hasOwnProperty("constraint"))
                    object.constraint = message.constraint;
                if (message.assertionValue != null && message.hasOwnProperty("assertionValue")) {
                    object.assertionValue = message.assertionValue;
                    if (options.oneofs)
                        object._assertionValue = "assertionValue";
                }
                return object;
            };
    
            /**
             * Converts this SMTConstraint to JSON.
             * @function toJSON
             * @memberof eidin.SMTConstraint
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            SMTConstraint.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };
    
            /**
             * Gets the default type url for SMTConstraint
             * @function getTypeUrl
             * @memberof eidin.SMTConstraint
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            SMTConstraint.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/eidin.SMTConstraint";
            };
    
            return SMTConstraint;
        })();
    
        eidin.CallbackId = (function() {
    
            /**
             * Properties of a CallbackId.
             * @memberof eidin
             * @interface ICallbackId
             * @property {number|Long|null} [bytesStart] CallbackId bytesStart
             * @property {number|Long|null} [bytesEnd] CallbackId bytesEnd
             */
    
            /**
             * Constructs a new CallbackId.
             * @memberof eidin
             * @classdesc Represents a CallbackId.
             * @implements ICallbackId
             * @constructor
             * @param {eidin.ICallbackId=} [properties] Properties to set
             */
            function CallbackId(properties) {
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }
    
            /**
             * CallbackId bytesStart.
             * @member {number|Long} bytesStart
             * @memberof eidin.CallbackId
             * @instance
             */
            CallbackId.prototype.bytesStart = $util.Long ? $util.Long.fromBits(0,0,true) : 0;
    
            /**
             * CallbackId bytesEnd.
             * @member {number|Long} bytesEnd
             * @memberof eidin.CallbackId
             * @instance
             */
            CallbackId.prototype.bytesEnd = $util.Long ? $util.Long.fromBits(0,0,true) : 0;
    
            /**
             * Creates a new CallbackId instance using the specified properties.
             * @function create
             * @memberof eidin.CallbackId
             * @static
             * @param {eidin.ICallbackId=} [properties] Properties to set
             * @returns {eidin.CallbackId} CallbackId instance
             */
            CallbackId.create = function create(properties) {
                return new CallbackId(properties);
            };
    
            /**
             * Encodes the specified CallbackId message. Does not implicitly {@link eidin.CallbackId.verify|verify} messages.
             * @function encode
             * @memberof eidin.CallbackId
             * @static
             * @param {eidin.ICallbackId} message CallbackId message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            CallbackId.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.bytesStart != null && Object.hasOwnProperty.call(message, "bytesStart"))
                    writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.bytesStart);
                if (message.bytesEnd != null && Object.hasOwnProperty.call(message, "bytesEnd"))
                    writer.uint32(/* id 2, wireType 0 =*/16).uint64(message.bytesEnd);
                return writer;
            };
    
            /**
             * Encodes the specified CallbackId message, length delimited. Does not implicitly {@link eidin.CallbackId.verify|verify} messages.
             * @function encodeDelimited
             * @memberof eidin.CallbackId
             * @static
             * @param {eidin.ICallbackId} message CallbackId message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            CallbackId.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };
    
            /**
             * Decodes a CallbackId message from the specified reader or buffer.
             * @function decode
             * @memberof eidin.CallbackId
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {eidin.CallbackId} CallbackId
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            CallbackId.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.eidin.CallbackId();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1: {
                            message.bytesStart = reader.uint64();
                            break;
                        }
                    case 2: {
                            message.bytesEnd = reader.uint64();
                            break;
                        }
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };
    
            /**
             * Decodes a CallbackId message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof eidin.CallbackId
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {eidin.CallbackId} CallbackId
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            CallbackId.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };
    
            /**
             * Verifies a CallbackId message.
             * @function verify
             * @memberof eidin.CallbackId
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            CallbackId.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.bytesStart != null && message.hasOwnProperty("bytesStart"))
                    if (!$util.isInteger(message.bytesStart) && !(message.bytesStart && $util.isInteger(message.bytesStart.low) && $util.isInteger(message.bytesStart.high)))
                        return "bytesStart: integer|Long expected";
                if (message.bytesEnd != null && message.hasOwnProperty("bytesEnd"))
                    if (!$util.isInteger(message.bytesEnd) && !(message.bytesEnd && $util.isInteger(message.bytesEnd.low) && $util.isInteger(message.bytesEnd.high)))
                        return "bytesEnd: integer|Long expected";
                return null;
            };
    
            /**
             * Creates a CallbackId message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof eidin.CallbackId
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {eidin.CallbackId} CallbackId
             */
            CallbackId.fromObject = function fromObject(object) {
                if (object instanceof $root.eidin.CallbackId)
                    return object;
                var message = new $root.eidin.CallbackId();
                if (object.bytesStart != null)
                    if ($util.Long)
                        (message.bytesStart = $util.Long.fromValue(object.bytesStart)).unsigned = true;
                    else if (typeof object.bytesStart === "string")
                        message.bytesStart = parseInt(object.bytesStart, 10);
                    else if (typeof object.bytesStart === "number")
                        message.bytesStart = object.bytesStart;
                    else if (typeof object.bytesStart === "object")
                        message.bytesStart = new $util.LongBits(object.bytesStart.low >>> 0, object.bytesStart.high >>> 0).toNumber(true);
                if (object.bytesEnd != null)
                    if ($util.Long)
                        (message.bytesEnd = $util.Long.fromValue(object.bytesEnd)).unsigned = true;
                    else if (typeof object.bytesEnd === "string")
                        message.bytesEnd = parseInt(object.bytesEnd, 10);
                    else if (typeof object.bytesEnd === "number")
                        message.bytesEnd = object.bytesEnd;
                    else if (typeof object.bytesEnd === "object")
                        message.bytesEnd = new $util.LongBits(object.bytesEnd.low >>> 0, object.bytesEnd.high >>> 0).toNumber(true);
                return message;
            };
    
            /**
             * Creates a plain object from a CallbackId message. Also converts values to other types if specified.
             * @function toObject
             * @memberof eidin.CallbackId
             * @static
             * @param {eidin.CallbackId} message CallbackId
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            CallbackId.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                var object = {};
                if (options.defaults) {
                    if ($util.Long) {
                        var long = new $util.Long(0, 0, true);
                        object.bytesStart = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.bytesStart = options.longs === String ? "0" : 0;
                    if ($util.Long) {
                        var long = new $util.Long(0, 0, true);
                        object.bytesEnd = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.bytesEnd = options.longs === String ? "0" : 0;
                }
                if (message.bytesStart != null && message.hasOwnProperty("bytesStart"))
                    if (typeof message.bytesStart === "number")
                        object.bytesStart = options.longs === String ? String(message.bytesStart) : message.bytesStart;
                    else
                        object.bytesStart = options.longs === String ? $util.Long.prototype.toString.call(message.bytesStart) : options.longs === Number ? new $util.LongBits(message.bytesStart.low >>> 0, message.bytesStart.high >>> 0).toNumber(true) : message.bytesStart;
                if (message.bytesEnd != null && message.hasOwnProperty("bytesEnd"))
                    if (typeof message.bytesEnd === "number")
                        object.bytesEnd = options.longs === String ? String(message.bytesEnd) : message.bytesEnd;
                    else
                        object.bytesEnd = options.longs === String ? $util.Long.prototype.toString.call(message.bytesEnd) : options.longs === Number ? new $util.LongBits(message.bytesEnd.low >>> 0, message.bytesEnd.high >>> 0).toNumber(true) : message.bytesEnd;
                return object;
            };
    
            /**
             * Converts this CallbackId to JSON.
             * @function toJSON
             * @memberof eidin.CallbackId
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            CallbackId.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };
    
            /**
             * Gets the default type url for CallbackId
             * @function getTypeUrl
             * @memberof eidin.CallbackId
             * @static
             * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns {string} The default type url
             */
            CallbackId.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                if (typeUrlPrefix === undefined) {
                    typeUrlPrefix = "type.googleapis.com";
                }
                return typeUrlPrefix + "/eidin.CallbackId";
            };
    
            return CallbackId;
        })();
    
        return eidin;
    })();

    return $root;
});
