// List of Domain Creating Queries
CREATE (d1:DOMAIN{
	domainName:'Networking',
    subDomains:['Network Architecture','Network Security','Wireless Communication', 'Routing & Switching']
});
CREATE (d2:DOMAIN{
	domainName:'Cloud Computing',
    subDomains:["Infrastructure as a Service (IaaS)",
    "Platform as a Service (PaaS)",
    "Cloud Security",
    "Serverless Computing"]
})
CREATE (d3:DOMAIN{
	domainName:'Software Development',
    subDomains:[ "Frontend Development",
    "Backend Development",
    "Mobile App Development",
    "Software Testing & QA"]
})
CREATE (d4:DOMAIN{
	domainName:'Cybersecurity',
    subDomains:["Application Security",
    "Network Security",
    "Identity and Access Management (IAM)",
    "Security Operations (SecOps)"]
})
CREATE (d5:DOMAIN{
	domainName:'Artificial Intelligence and Machine Learning',
    subDomains:[ "Natural Language Processing (NLP)",
    "Computer Vision",
    "Deep Learning",
    "Reinforcement Learning"]
})
CREATE (d6:DOMAIN{
	domainName:'DevOps',
    subDomains:[ "CI/CD Pipelines",
    "Infrastructure as Code (IaC)",
    "Monitoring & Logging",
    "Containerization & Orchestration"]
})
CREATE (d7:DOMAIN{
	domainName:'Internet of Things',
    subDomains:[ "Embedded Systems",
    "IoT Protocols",
    "Edge Computing",
    "IoT Security"]
})
CREATE (d8:DOMAIN{
	domainName:'Blockchain',
    subDomains:[ "Smart Contracts",
    "Consensus Mechanisms",
    "DeFi (Decentralized Finance)",
    "NFTs & Tokenization"]
})

// List of Tool Creating Queries
// Model
CREATE (t-placeHolder:TOOL{
    toolName:'T_N_Architecture-placeHolder'
    likes:0
    communicationCount:0
    deliveryFormat:''
})
// Program for getting Tools created
//  StringBuilder sb = new StringBuilder();
//        int[] toolsCount=new int[]{50,70,30,40,20,60,20,40};
//        List<String[]> subDomains=new ArrayList<String[]>();
//        subDomains.add(new String[]{"Architecture","Security","Communication","Routing"});
//        subDomains.add(new String[]{"Infrastructure","Platform","Cloud","Serverless"});
//        subDomains.add(new String[]{"Frontend","Backend","Mobile App","Software Testing"});
//        subDomains.add(new String[]{"Application","Network","Identity","Operations"});
//        subDomains.add(new String[]{"NLP","Computer Vision","Deep Learning","Reinforcement"});
//        subDomains.add(new String[]{"CI/CD Pipelines","IaC","Monitoring","Containerization"});
//        subDomains.add(new String[]{"Embedded","Protocols","Computing","IoT Security"});
//        subDomains.add(new String[]{"Contracts","Consensus","DeFi","Tokenization"});
//        int toolsCountArrayIndex=0;
//        int toolGeneratorPlaceHolder=1;
//        for(String[] subDomain:subDomains){
//                 int requiredCountToolTogenerate=toolsCount[toolsCountArrayIndex];
//                for(String j:subDomain){
//                 sb.setLength(0);
//                 for (int i = 1; i <= requiredCountToolTogenerate; i++) {
//                     sb.append("CREATE (t")
//                       .append(toolGeneratorPlaceHolder)
//                       .append(":TOOL {\n")
//                       .append("    toolName: 'T_N_").append(j).append('-')
//                       .append(i)
//                       .append("',\n")
//                       .append("    likes: 0,\n")
//                       .append("    communicationCount: 0,\n")
//                       .append("    deliveryFormat: ''\n")
//                       .append("})\n");
//                     toolGeneratorPlaceHolder=toolGeneratorPlaceHolder+1;
//                 }
//                 System.out.println(sb.toString());
//             }
//             toolsCountArrayIndex=toolsCountArrayIndex+1;
//        }

// 
0,1,2,3,4,5,6,7,8,9
