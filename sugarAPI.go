package main

mergeGroup := map[string] map{
'conv'  = > array(
'' = > 'www.altoros.com',
'altoros.com'  = > 'www.altoros.com',
'www.altoros.com' = > 'www.altoros.com',
'altoroslabs'  = > 'labs.altoros.com',
'www.altoroslabs'  = > 'labs.altoros.com',
'altoroslabs.com'  = > 'labs.altoros.com',
'www.altoroslabs.com' = > 'labs.altoros.com',
'altorosno'  = > 'www.altoros.no',
'altoros.no'  = > 'www.altoros.no',
'www.altoros.no'  = > 'www.altoros.no',
'altorosdk'  = > 'www.altoros.dk',
'altoros.dk'  = > 'www.altoros.dk',
'www.altoros.dk' = > 'www.altoros.dk',
'altorosde'  = > 'www.altoros.de',
'altoros.de'  = > 'www.altoros.de',
'www.altoros.de'  = > 'www.altoros.de',
'akelioscom'  = > 'www.akelios.com',
'akelios.com'  = > 'www.akelios.com',
'www.akelios.com' = > 'www.akelios.com',
'altorosbigdatahub.com'  = > 'www.altorosbigdatahub.com',
'www.altorosbigdatahub.com'  = > 'www.altorosbigdatahub.com',
'teamexpand'  = > 'www.teamexpand.com',
'teamexpand.com' = > 'www.teamexpand.com',
'www.teamexpand.com'  = > 'www.teamexpand.com',
'apatar' = > 'www.apatar.com',
'apatar.com'  = > 'www.apatar.com',
'www.apatar.com'  = > 'www.apatar.com',
'rubyonrailsteam'  = > 'www.rubyonrailsteam.com',
'rubyonrailsteam.com' = > 'www.rubyonrailsteam.com',
'www.rubyonrailsteam.com'  = > 'www.rubyonrailsteam.com',
'fi.altoros.com'  = > 'www.altoros.fi',
'altoros.fi'  => 'www.altoros.fi',
'www.altoros.fi'  = > 'www.altoros.fi',
),
'wps'  = > array(
'' = > 'www.altoros.com',
'altoros.com'  = > 'www.altoros.com',
'www.altoros.com'  = > 'www.altoros.com',
'altoroslabs'  = > 'labs.altoros.com',
'www.altoroslabs'  = > 'labs.altoros.com',
'altoroslabs.com'  = > 'labs.altoros.com',
'www.altoroslabs.com'  = > 'labs.altoros.com',
'altorosno'  = > 'www.altoros.no',
'altoros.no'  = > 'www.altoros.no',
'www.altoros.no'  = > 'www.altoros.no',
'altorosdk'  = > 'www.altoros.dk',
'altoros.dk'  = > 'www.altoros.dk',
'www.altoros.dk'  = > 'www.altoros.dk',
'altorosde'  = > 'www.altoros.de',
'altoros.de'  = > 'www.altoros.de',
'www.altoros.de'  = > 'www.altoros.de',
'akelioscom'  = > 'www.akelios.com',
'akelios.com'  = > 'www.akelios.com',
'www.akelios.com'  = > 'www.akelios.com',
'altorosbigdatahub.com'  = > 'www.altorosbigdatahub.com',
'www.altorosbigdatahub.com' = > 'www.altorosbigdatahub.com',
'teamexpand'  = > 'www.teamexpand.com',
'teamexpand.com'  = > 'www.teamexpand.com',
'www.teamexpand.com'  = > 'www.teamexpand.com',
'apatar'  = > 'www.apatar.com',
'apatar.com'  = > 'www.apatar.com',
'www.apatar.com' = > 'www.apatar.com',
'rubyonrailsteam'  = > 'www.rubyonrailsteam.com',
'rubyonrailsteam.com'  = > 'www.rubyonrailsteam.com',
'www.rubyonrailsteam.com'  = > 'www.rubyonrailsteam.com',
'fi.altoros.com'  = > 'www.altoros.fi',
'altoros.fi'  = > 'www.altoros.fi',
'www.altoros.fi'  = > 'www.altoros.fi',
),
'opps'  => array(
'' = > 'www.altoros.com',
'altoros.com'  => 'www.altoros.com',
'www.altoros.com'  = > 'www.altoros.com',
'altoroslabs'  => 'labs.altoros.com',
'www.altoroslabs'  = > 'labs.altoros.com',
'altoroslabs.com'  => 'labs.altoros.com',
'www.altoroslabs.com'  = > 'labs.altoros.com',
'altorosno'  => 'www.altoros.no',
'altoros.no'  = > 'www.altoros.no',
'www.altoros.no'  = > 'www.altoros.no',
'altorosdk'  = > 'www.altoros.dk',
'altoros.dk'  => 'www.altoros.dk',
'www.altoros.dk'  = > 'www.altoros.dk',
'altorosde'  => 'www.altoros.de',
'altoros.de'  = > 'www.altoros.de',
'www.altoros.de'  = > 'www.altoros.de',
'akelioscom'  = > 'www.akelios.com',
'akelios.com'  => 'www.akelios.com',
'www.akelios.com'  = > 'www.akelios.com',
'altorosbigdatahub.com'  = > 'www.altorosbigdatahub.com',
'www.altorosbigdatahub.com'  = > 'www.altorosbigdatahub.com',
'teamexpand'  = > 'www.teamexpand.com',
'teamexpand.com'  = > 'www.teamexpand.com',
'www.teamexpand.com'  = > 'www.teamexpand.com',
'apatar'  = > 'www.apatar.com',
'apatar.com'  = > 'www.apatar.com',
'www.apatar.com'  = > 'www.apatar.com',
'rubyonrailsteam'  = > 'www.rubyonrailsteam.com',
'rubyonrailsteam.com'  = > 'www.rubyonrailsteam.com',
'www.rubyonrailsteam.com'  = > 'www.rubyonrailsteam.com',
'fi.altoros.com' = > 'www.altoros.fi',
'altoros.fi'  = > 'www.altoros.fi',
'www.altoros.fi'  = > 'www.altoros.fi',
),
'wps_name'  = > array(
'Re: Hadoop Distributions: Cloudera vs. Hortonworks vs.MapR'  = > 'Hadoop Distributions: Evaluating Cloudera, Hortonworks, and MapR in Micro-benchmarks and Real-world Applications',
'Re: Hadoop Distributions: Evaluating Cloudera, Hortonworks, and MapR in Micro-benchmarks and Real-world Applications'  = > 'Hadoop Distributions: Evaluating Cloudera, Hortonworks, and MapR in Micro-benchmarks and Real-world Applications',
'Re: Hadoop Distributions: Evaluating Cloudera Hortonworks and MapR in Micro-benchmarks and Real-world Applications'  = > 'Hadoop Distributions: Evaluating Cloudera, Hortonworks, and MapR in Micro-benchmarks and Real-world Applications',
'Re: Cloud Platform Comparison - CloudStack, Eucalyptus, vCloud Director, and Openstack'  = > 'Cloud Platform Comparison - CloudStack, Eucalyptus, vCloud Director, and Openstack',
'Re: Cloud Platform Comparison - CloudStack Eucalyptus vCloud Director and Openstack' = > 'Cloud Platform Comparison - CloudStack, Eucalyptus, vCloud Director, and Openstack',
'Re: A Vendor-independent Comparison of NoSQL Databases: Cassandra, HBase, MongoDB, Riak'  = > 'A Vendor-independent Comparison of NoSQL Databases: Cassandra, HBase, MongoDB, Riak',
'Re: A Vendor-independent Comparison of NoSQL Databases Cassandra HBase MongoDB Riak'  = > 'A Vendor-independent Comparison of NoSQL Databases: Cassandra, HBase, MongoDB, Riak',
'Re: A Guide on Installing the OpenStack and Cloud Foundry PaaS on HP Moonshot'  = > 'A Guide on Installing the OpenStack and Cloud Foundry PaaS on HP Moonshot',
'A Guide on Installing the OpenStack and Cloud Foundry PaaS on HP Moonshot'  = > 'A Guide on Installing the OpenStack and Cloud Foundry PaaS on HP Moonshot',
'Re: Five Ways to Measure Your Programmers\' Performance'  = > 'Five Ways to Measure Your Programmers\' Performance',
'Re: Five Ways to Measure Your Programmers Performance'  = > 'Five Ways to Measure Your Programmers\' Performance',
'Re: OpenShift and Cloud Foundry PaaS: High-level Overview of Features and Architectures'  => 'OpenShift and Cloud Foundry PaaS: High-level Overview of Features and Architectures',
'Re: Data Visualization Frameworks: Flot vs.Highcharts vs.D3'  = > 'Data Visualization Frameworks: Flot vs.Highcharts vs.D3',
'Re: Data Visualization Frameworks: Flot vs.Highcharts vs.D3.js'  = > 'Data Visualization Frameworks: Flot vs.Highcharts vs.D3',
'Re: Data Visualization Frameworks: Flot vs.Highcharts vs.D3.s'  => 'Data Visualization Frameworks: Flot vs.Highcharts vs.D3',
'Re:'  => '',
'Re: Hadoop-based Movie Recommendation Engine: A Comparison of the Apriori Algorithm vs.the k-means Method'  = > 'Hadoop-based Movie Recommendation Engine: A Comparison of the Apriori Algorithm vs.the k-means Method',
'Re: Hadoop on Windows Azure: Hive vs.JavaScript for Processing Big Data'  => 'Hadoop on Windows Azure: Hive vs. JavaScript for Processing Big Data',
'Re: Benchmarking Couchbase Server for Interactive Applications'  = > 'Benchmarking Couchbase Server for Interactive Applications',
'Re: Software Product Outsourcing Guide'  = > 'Software Product Outsourcing Guide',
'Re: Boost Project Management Efficiency with Microsoft TFS 11'  = > 'Boost Project Management Efficiency with Microsoft TFS 11',
'webinar: Boost Project Management Efficiency with Microsoft TFS 11'  = > 'Boost Project Management Efficiency with Microsoft TFS 11',
'Re: Cloud Computing: How to Solve Challenges and Avoid Typical Mistakes Using Amazon Web Services'  = > 'Cloud Computing: How to Solve Challenges and Avoid Typical Mistakes Using Amazon Web Services',
'Re: Five Mistakes to Avoid in Creating User Interfaces for Mobile Devices'  => 'Five Mistakes to Avoid in Creating User Interfaces for Mobile Devices',
'Re: Comparing Ruby on Rails Public vs.Private Cloud Options'  => 'Comparing Ruby on Rails Public vs.Private Cloud Options',
'Re: How to Integrate Independent QA Testing to Shorten Development Cycles'  = > 'How to Integrate Independent QA Testing to Shorten Development Cycles',
'Re: Remote DBA Program 6-Step Guide and Most Common Mistakes'  = > 'Remote DBA Program 6-Step Guide and Most Common Mistakes',
'Re: How to Identify and Mitigate Security and Intellectual Property Risks When Outsourcing Offshore'  = > 'How to Identify and Mitigate Security and Intellectual Property Risks When Outsourcing Offshore',
'slideshare: How to Identify and Mitigate Security and Intellectual Property Risks When Outsourcing Offshore'  = > 'How to Identify and Mitigate Security and Intellectual Property Risks When Outsourcing Offshore',
'Re: How to Optimize ROI Using Remote DBA and Avoid the Most Common Mistakes'  = > 'How to Optimize ROI Using Remote DBA and Avoid the Most Common Mistakes',
'Re: Cut Costs with Adobe AIR Cross-platform App Development'  = > 'Cut Costs with Adobe AIR Cross-platform App Development',
),
}