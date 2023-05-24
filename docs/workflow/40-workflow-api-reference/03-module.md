---
toc_max_heading_level: 6
---

# Module Reference

### Module (top level object)

To create a module, a public git repository must be created with the `module.yaml` file.

A module is basically a group of steps.

The `module.yaml` file first goes through a templating engine before being parsed. So some variables are available:

- [{{ .Job }}](job#job-top-level-object) and its childs, which represent the Job object using the module. Can be useful if you want to dynamically set an value based on the job.
- [{{ .Step }}](job#steps-step) and its childs, which represent the Step object using the module. Can be useful if you want the step name.

Notice that the templating follows the Go format. You can also apply [sprig](http://masterminds.github.io/sprig/) templating functions.

To outputs environment variables, just append KEY=value to the "${DEEPSQUARE_ENV}" file.

An example of module is available [here](https://github.com/deepsquare-io/workflow-module-example).

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>name</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

Name of the module.

Go name: "Name".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>description</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

Description of the module.

Go name: "Description".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>minimumResources</strong></td>
<td valign="top"><a href="job#resources-jobresources">JobResources</a>!</td>
<td>

Minimum job resources.

Go name: "MinimumResources".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>inputs</strong></td>
<td valign="top">[<a href="#inputs-moduleinput">ModuleInput</a>!]</td>
<td>

List of allowed arguments.

Go name: "Inputs".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>outputs</strong></td>
<td valign="top">[<a href="#outputs-moduleoutput">ModuleOutput</a>!]</td>
<td>

List of exported environment variables.

Go name: "Outputs".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>steps</strong></td>
<td valign="top">[<a href="job#steps-step">Step</a>!]!</td>
<td>

Steps of the module.

Go name: "Steps".

</td>
</tr>
</tbody>
</table>

### `.inputs[]` _ModuleInput_

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>key</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

Name of the input.

Go name: "Key".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>description</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

Description of the input.

Go name: "Description".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>default</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

Default value.

If not set, will default to empty string.

Go name: "Default".

</td>
</tr>
</tbody>
</table>

### `.outputs[]` _ModuleOutput_

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>key</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

Name of the output.

Go name: "Key".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>description</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

Description of the output.

Go name: "Description".

</td>
</tr>
</tbody>
</table>

## Scalars

### Boolean

The `Boolean` scalar type represents `true` or `false`.

### Int

The `Int` scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1.

### String

The `String` scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used to represent free-form human-readable text.
