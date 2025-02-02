// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.Inputs
{

    /// <summary>
    /// CNTK (aka Microsoft Cognitive Toolkit) job settings.
    /// </summary>
    public sealed class CNTKsettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Command line arguments that need to be passed to the python script or cntk executable.
        /// </summary>
        [Input("commandLineArgs")]
        public Input<string>? CommandLineArgs { get; set; }

        /// <summary>
        /// Specifies the path of the BrainScript config file. This property can be specified only if the languageType is 'BrainScript'.
        /// </summary>
        [Input("configFilePath")]
        public Input<string>? ConfigFilePath { get; set; }

        /// <summary>
        /// The language to use for launching CNTK (aka Microsoft Cognitive Toolkit) job. Valid values are 'BrainScript' or 'Python'.
        /// </summary>
        [Input("languageType")]
        public Input<string>? LanguageType { get; set; }

        /// <summary>
        /// Number of processes to launch for the job execution. The default value for this property is equal to nodeCount property
        /// </summary>
        [Input("processCount")]
        public Input<int>? ProcessCount { get; set; }

        /// <summary>
        /// The path to the Python interpreter. This property can be specified only if the languageType is 'Python'.
        /// </summary>
        [Input("pythonInterpreterPath")]
        public Input<string>? PythonInterpreterPath { get; set; }

        /// <summary>
        /// Python script to execute. This property can be specified only if the languageType is 'Python'.
        /// </summary>
        [Input("pythonScriptFilePath")]
        public Input<string>? PythonScriptFilePath { get; set; }

        public CNTKsettingsArgs()
        {
        }
    }
}
