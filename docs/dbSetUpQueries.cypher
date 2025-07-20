// Code for Creating Domain, Subdomain and Relation between them
// import java.util.*;
// class Main {
//     public static void main(String[] args) {
//         StringBuilder sb = new StringBuilder();

//         // List of Domains
//         String[] domains=new String[]{"Networking",
//                                     "Cloud Computing",
//                                     "Software Development",
//                                     "Cybersecurity",
//                                     "Artificial Intelligence and Machine Learning",
//                                     "DevOps",
//                                     "Internet of Things",
//                                     "Blockchain"};

//         // For Creating Domain Nodes
//         for (int i = 0; i < domains.length; i++) {
//             sb.append("CREATE (d")
//               .append(i)
//               .append(":DOMAIN {\n")
//               .append("    domainName:'").append(domains[i]).append("'")
//               .append("})\n");
//             System.out.println(sb.toString());
//             sb.setLength(0);
//         }

//         // For Creating SubDomain Nodes
//         // List of Subdomains
//         List<String[]> subDomains=new ArrayList<String[]>();
//         subDomains.add(new String[]{"Architecture","Security","Communication","Routing"});
//         subDomains.add(new String[]{"Infrastructure","Platform","Cloud","Serverless"});
//         subDomains.add(new String[]{"Frontend","Backend","Mobile App","Software Testing"});
//         subDomains.add(new String[]{"Application","Network","Identity","Operations"});
//         subDomains.add(new String[]{"NLP","Computer Vision","Deep Learning","Reinforcement"});
//         subDomains.add(new String[]{"CI/CD Pipelines","IaC","Monitoring","Containerization"});
//         subDomains.add(new String[]{"Embedded","Protocols","Computing","IoT Security"});
//         subDomains.add(new String[]{"Contracts","Consensus","DeFi","Tokenization"});

//         int sdIndex=0;

//         for (int i = 0; i < subDomains.size(); i++) {
//             String[] str=subDomains.get(i);
//             for(String s:str){
//                 sb.append("CREATE (sd")
//                   .append(sdIndex)
//                   .append(":SUBDOMAIN {\n")
//                   .append("    subDomainName:'").append(s).append("'")
//                   .append("})\n");

//                 sb.append("MERGE (sd")
//                 .append(sdIndex).append(")").append("-[:PART_OF]->").append("(")
//                 .append('d').append(i).append(")");
//                 System.out.println(sb.toString());
//                 sb.setLength(0);
//                 sdIndex=sdIndex+1;
//             }
//         }
//     }
// }

// List of Tool Creating Queries
// Model
CREATE (t-placeHolder:TOOL{
  toolName:'T_N_Architecture-placeHolder'
  likes:0
  communicationCount:0
  deliveryFormat:''
  })
  
// Code - 2
// import java.util.*;
// class Main {
//     public static void main(String[] args) {
//         StringBuilder sb = new StringBuilder();
  
//         // List of Domains
//         String[] domains=new String[]{"Networking",
//                                     "Cloud Computing",
//                                     "Software Development",
//                                     "Cybersecurity",
//                                     "Artificial Intelligence and Machine Learning",
//                                     "DevOps",
//                                     "Internet of Things",
//                                     "Blockchain"};
  
//         HashMap<String,Integer> countForCreation= new HashMap<String,Integer>();
  
//         countForCreation.put("Networking-Architecture",12);
//         countForCreation.put("Networking-Security",10);
//         countForCreation.put("Networking-Communication",15);
//         countForCreation.put("Networking-Routing",20);
//         countForCreation.put("Cloud Computing-Infrastructure",21);
//         countForCreation.put("Cloud Computing-Platform",10);
//         countForCreation.put("Cloud Computing-Cloud",30);
//         countForCreation.put("Cloud Computing-Serverless",22);
//         countForCreation.put("Software Development-Frontend",15);
//         countForCreation.put("Software Development-Backend",21);
//         countForCreation.put("Software Development-Mobile App",25);
//         countForCreation.put("Software Development-Software Testing",8);
//         countForCreation.put("Cybersecurity-Application",5);
//         countForCreation.put("Cybersecurity-Network",8);
//         countForCreation.put("Cybersecurity-Identity",24);
//         countForCreation.put("Cybersecurity-Operations",7);
//         countForCreation.put("Artificial Intelligence and Machine Learning-NLP",3);
//         countForCreation.put("Artificial Intelligence and Machine Learning-Computer Vision",5);
//         countForCreation.put("Artificial Intelligence and Machine Learning-Deep Learning",7);
//         countForCreation.put("Artificial Intelligence and Machine Learning-Reinforcement",2);
//         countForCreation.put("DevOps-CI/CD Pipelines",33);
//         countForCreation.put("DevOps-IaC",10);
//         countForCreation.put("DevOps-Monitoring",11);
//         countForCreation.put("DevOps-Containerization",7);
//         countForCreation.put("Internet of Things-Embedded",10);
//         countForCreation.put("Internet of Things-Protocols",45);
//         countForCreation.put("Internet of Things-Computing",20);
//         countForCreation.put("Internet of Things-IoT Security",8);
//         countForCreation.put("Blockchain-Contracts",11);
//         countForCreation.put("Blockchain-Consensus",16);
//         countForCreation.put("Blockchain-DeFi",18);
//         countForCreation.put("Blockchain-Tokenization",40);
  
//         // List of Subdomains
//         List<String[]> subDomains=new ArrayList<String[]>();
//         subDomains.add(new String[]{"Architecture","Security","Communication","Routing"});
//         subDomains.add(new String[]{"Infrastructure","Platform","Cloud","Serverless"});
//         subDomains.add(new String[]{"Frontend","Backend","Mobile App","Software Testing"});
//         subDomains.add(new String[]{"Application","Network","Identity","Operations"});
//         subDomains.add(new String[]{"NLP","Computer Vision","Deep Learning","Reinforcement"});
//         subDomains.add(new String[]{"CI/CD Pipelines","IaC","Monitoring","Containerization"});
//         subDomains.add(new String[]{"Embedded","Protocols","Computing","IoT Security"});
//         subDomains.add(new String[]{"Contracts","Consensus","DeFi","Tokenization"});
  
//         // Index for Outer for Loop
//         int arrIndex=0;
//         // Index for query Execution
//         int toolGeneratorPlaceHolder=1;
//         for(String[] subDomain:subDomains){
//             String domain=domains[arrIndex];
//             StringBuilder key=new StringBuilder();
//             for(String j:subDomain){
//                 key.setLength(0);
//                 key.append(domain).append("-").append(j);
//                 int requiredCountToolTogenerate=countForCreation.get(key.toString());
//                 sb.setLength(0);
//                 for (int i = 1; i <= requiredCountToolTogenerate; i++) {
//                     sb.append("CREATE (t")
//                       .append(toolGeneratorPlaceHolder)
//                       .append(":TOOL {\n")
//                       .append("    toolName: 'T_N_").append(domain.substring(0,3)).append(j.substring(0,3)).append('-')
//                       .append(i)
//                       .append("',\n")
//                       .append("    likes: 0,\n")
//                       .append("    communicationCount: 0,\n")
//                       .append("    deliveryFormat: ''\n")
//                       .append("})\n");
//                     toolGeneratorPlaceHolder=toolGeneratorPlaceHolder+1;
//                 }
//                 System.out.println(sb.toString());
//         }
//         arrIndex=arrIndex+1;
//         }
//     }
// }
  
// Code - 3
// Code for Connecting Tool and SubDomain
// Query
  MATCH (d31:DOMAIN { domainName: 'Blockchain' })
  MATCH (sd31:SUBDOMAIN{ subDomainName:'Tokenization' })-[:PART_OF]->(d31)
  MATCH (t31:TOOL)
  WHERE t31.toolName CONTAINS 'BloTok'
  MERGE (t31)-[:USED_FOR]->(sd31)
// import java.util.*;
// class Main {
//     public static void main(String[] args) {
//         StringBuilder sb = new StringBuilder();
  
//         // List of Domains
//         String[] domains=new String[]{"Networking",
//                                     "Cloud Computing",
//                                     "Software Development",
//                                     "Cybersecurity",
//                                     "Artificial Intelligence and Machine Learning",
//                                     "DevOps",
//                                     "Internet of Things",
//                                     "Blockchain"};
  
//         // List of Subdomains
//         List<String[]> subDomains=new ArrayList<String[]>();
//         subDomains.add(new String[]{"Architecture","Security","Communication","Routing"});
//         subDomains.add(new String[]{"Infrastructure","Platform","Cloud","Serverless"});
//         subDomains.add(new String[]{"Frontend","Backend","Mobile App","Software Testing"});
//         subDomains.add(new String[]{"Application","Network","Identity","Operations"});
//         subDomains.add(new String[]{"NLP","Computer Vision","Deep Learning","Reinforcement"});
//         subDomains.add(new String[]{"CI/CD Pipelines","IaC","Monitoring","Containerization"});
//         subDomains.add(new String[]{"Embedded","Protocols","Computing","IoT Security"});
//         subDomains.add(new String[]{"Contracts","Consensus","DeFi","Tokenization"});
//         int queryIndex=0;
//         for(int i=0;i<domains.length;i++){
//             String[] subDomainList=subDomains.get(i);
//             for(String str:subDomainList){
//                 sb.append(" MATCH (d").append(queryIndex).append(":DOMAIN {domainName: '").append(domains[i])
//                 .append("'})\n")
//                 .append(" MATCH (sd").append(queryIndex).append(":SUBDOMAIN").append("{subDomainName:'").append(str)
//                 .append("'})-[:PART_OF]->(d").append(queryIndex)
//                 .append(")\nMATCH (t").append(queryIndex).append(":TOOL)\n")
//                 .append("WHERE t").append(queryIndex).append(".toolName contains '").append(domains[i].substring(0,3))
//                 .append(str.substring(0,3))
//                 .append("'\nMERGE (t")
//                 .append(queryIndex)
//                 .append(")-[:USED_FOR]->(sd")
//                 .append(queryIndex)
//                 .append(") \n WITH 1 AS dummy \n");
//                 System.out.print(sb.toString());
//                 sb.setLength(0);
//                 queryIndex=queryIndex+1;
//             }
//         }
//     }
// }
  
//  Assign DeliveryFormat
//   0, 1 - Open Source
//   2, 3, 8 - Package Manager
//   4, 5, 9 - Web Based
//   6, 7 - IDE Plugin
// MATCH (N:TOOL)
// WHERE N.toolName ENDS WITH '2'
// SET N.deliveryFormat='Package Source' | 'Open Source' | 'Web Based' | 'IDE Plugin'
// RETURN N
  
//
  
// Program for assing the commuicationCount and likes for tool node
// import java.util.*;
// import java.util.Random;
// class Main {
//     public static void main(String[] args) {
//         StringBuilder sb = new StringBuilder();
//         Random random = new Random();
//         // List of Domains
//         String[] domains=new String[]{"Networking",
//                                     "Cloud Computing",
//                                     "Software Development",
//                                     "Cybersecurity",
//                                     "Artificial Intelligence and Machine Learning",
//                                     "DevOps",
//                                     "Internet of Things",
//                                     "Blockchain"};
  
//         HashMap<String,Integer> countForCreation= new HashMap<String,Integer>();
  
//         countForCreation.put("Networking-Architecture",12);
//         countForCreation.put("Networking-Security",10);
//         countForCreation.put("Networking-Communication",15);
//         countForCreation.put("Networking-Routing",20);
//         countForCreation.put("Cloud Computing-Infrastructure",21);
//         countForCreation.put("Cloud Computing-Platform",10);
//         countForCreation.put("Cloud Computing-Cloud",30);
//         countForCreation.put("Cloud Computing-Serverless",22);
//         countForCreation.put("Software Development-Frontend",15);
//         countForCreation.put("Software Development-Backend",21);
//         countForCreation.put("Software Development-Mobile App",25);
//         countForCreation.put("Software Development-Software Testing",8);
//         countForCreation.put("Cybersecurity-Application",5);
//         countForCreation.put("Cybersecurity-Network",8);
//         countForCreation.put("Cybersecurity-Identity",24);
//         countForCreation.put("Cybersecurity-Operations",7);
//         countForCreation.put("Artificial Intelligence and Machine Learning-NLP",3);
//         countForCreation.put("Artificial Intelligence and Machine Learning-Computer Vision",5);
//         countForCreation.put("Artificial Intelligence and Machine Learning-Deep Learning",7);
//         countForCreation.put("Artificial Intelligence and Machine Learning-Reinforcement",2);
//         countForCreation.put("DevOps-CI/CD Pipelines",33);
//         countForCreation.put("DevOps-IaC",10);
//         countForCreation.put("DevOps-Monitoring",11);
//         countForCreation.put("DevOps-Containerization",7);
//         countForCreation.put("Internet of Things-Embedded",10);
//         countForCreation.put("Internet of Things-Protocols",45);
//         countForCreation.put("Internet of Things-Computing",20);
//         countForCreation.put("Internet of Things-IoT Security",8);
//         countForCreation.put("Blockchain-Contracts",11);
//         countForCreation.put("Blockchain-Consensus",16);
//         countForCreation.put("Blockchain-DeFi",18);
//         countForCreation.put("Blockchain-Tokenization",40);
  
//         // List of Subdomains
//         List<String[]> subDomains=new ArrayList<String[]>();
//         subDomains.add(new String[]{"Architecture","Security","Communication","Routing"});
//         subDomains.add(new String[]{"Infrastructure","Platform","Cloud","Serverless"});
//         subDomains.add(new String[]{"Frontend","Backend","Mobile App","Software Testing"});
//         subDomains.add(new String[]{"Application","Network","Identity","Operations"});
//         subDomains.add(new String[]{"NLP","Computer Vision","Deep Learning","Reinforcement"});
//         subDomains.add(new String[]{"CI/CD Pipelines","IaC","Monitoring","Containerization"});
//         subDomains.add(new String[]{"Embedded","Protocols","Computing","IoT Security"});
//         subDomains.add(new String[]{"Contracts","Consensus","DeFi","Tokenization"});
  
//         // Index for Outer for Loop
//         int arrIndex=0;
//         // Index for query Execution
//         int toolGeneratorPlaceHolder=1;
//         for(String[] subDomain:subDomains){
//             String domain=domains[arrIndex];
//             StringBuilder key=new StringBuilder();
//             for(String j:subDomain){
//                 key.setLength(0);
//                 key.append(domain).append("-").append(j);
//                 int requiredCountToolTogenerate=countForCreation.get(key.toString());
//                 for (int i = 1; i <= requiredCountToolTogenerate; i++) {
//                      sb.setLength(0);
//                     int randomBounded = random.nextInt(30);
//                     sb.append("MATCH (t")
//                       .append(toolGeneratorPlaceHolder)
//                       .append(":TOOL {\n")
//                       .append("    toolName: 'T_N_").append(domain.substring(0,3)).append(j.substring(0,3)).append('-')
//                       .append(i)
//                       .append("'})\n SET t")
//                       .append(toolGeneratorPlaceHolder)
//                       .append(".communicationCount=")
//                       .append(randomBounded).append(";");
//                     toolGeneratorPlaceHolder=toolGeneratorPlaceHolder+1;
//                     System.out.println(sb.toString());
//                 }
  
//         }
//         arrIndex=arrIndex+1;
//         }
//     }
// }


// Run this query for all the nodes label, primary id as id are added to all nodes
// MATCH (t:Tool)
// WITH t
// ORDER BY id(t) ASC
// WITH collect(t) AS tools
// UNWIND range(0, size(tools)-1) AS idx
// WITH tools[idx] AS node, idx + 1 AS newId
// SET node.id = newId


// Online Java Compiler
// Use this editor to write, compile and run your Java code online

import java.util.*;
class Main {
    public static void main(String[] args) {
        StringBuilder sb = new StringBuilder();
         
        // List of Domains
        String[] domains=new String[]{"Networking",
                                    "Cloud Computing",
                                    "Software Development",
                                    "Cybersecurity",
                                    "Artificial Intelligence and Machine Learning",
                                    "DevOps",
                                    "Internet of Things",
                                    "Blockchain"};
        
        HashMap<String,Integer> countForCreation= new HashMap<String,Integer>();
        
        countForCreation.put("Networking-Architecture",12);
        countForCreation.put("Networking-Security",10);
        countForCreation.put("Networking-Communication",15);
        countForCreation.put("Networking-Routing",20);
        countForCreation.put("Cloud Computing-Infrastructure",21);
        countForCreation.put("Cloud Computing-Platform",10);
        countForCreation.put("Cloud Computing-Cloud",30);
        countForCreation.put("Cloud Computing-Serverless",22);
        countForCreation.put("Software Development-Frontend",15);
        countForCreation.put("Software Development-Backend",21);
        countForCreation.put("Software Development-Mobile App",25);
        countForCreation.put("Software Development-Software Testing",8);
        countForCreation.put("Cybersecurity-Application",5);
        countForCreation.put("Cybersecurity-Network",8);
        countForCreation.put("Cybersecurity-Identity",24);
        countForCreation.put("Cybersecurity-Operations",7);
        countForCreation.put("Artificial Intelligence and Machine Learning-NLP",3);
        countForCreation.put("Artificial Intelligence and Machine Learning-Computer Vision",5);
        countForCreation.put("Artificial Intelligence and Machine Learning-Deep Learning",7);
        countForCreation.put("Artificial Intelligence and Machine Learning-Reinforcement",2);
        countForCreation.put("DevOps-CI/CD Pipelines",33);
        countForCreation.put("DevOps-IaC",10);
        countForCreation.put("DevOps-Monitoring",11);
        countForCreation.put("DevOps-Containerization",7);
        countForCreation.put("Internet of Things-Embedded",10);
        countForCreation.put("Internet of Things-Protocols",45);
        countForCreation.put("Internet of Things-Computing",20);
        countForCreation.put("Internet of Things-IoT Security",8);
        countForCreation.put("Blockchain-Contracts",11);
        countForCreation.put("Blockchain-Consensus",16);
        countForCreation.put("Blockchain-DeFi",18);
        countForCreation.put("Blockchain-Tokenization",40);


        
        // List of Subdomains
        List<String[]> subDomains=new ArrayList<String[]>();
        subDomains.add(new String[]{"Architecture","Security","Communication","Routing"});
        subDomains.add(new String[]{"Infrastructure","Platform","Cloud","Serverless"});
        subDomains.add(new String[]{"Frontend","Backend","Mobile App","Software Testing"});
        subDomains.add(new String[]{"Application","Network","Identity","Operations"});
        subDomains.add(new String[]{"NLP","Computer Vision","Deep Learning","Reinforcement"});
        subDomains.add(new String[]{"CI/CD Pipelines","IaC","Monitoring","Containerization"});
        subDomains.add(new String[]{"Embedded","Protocols","Computing","IoT Security"});
        subDomains.add(new String[]{"Contracts","Consensus","DeFi","Tokenization"});
        
        // Index for Outer for Loop
        int arrIndex=0;
        // Index for query Execution
        int toolGeneratorPlaceHolder=1;
        for(String[] subDomain:subDomains){
            String domain=domains[arrIndex];
            StringBuilder key=new StringBuilder();
            for(String j:subDomain){
                key.setLength(0);
                key.append(domain).append("-").append(j);
                int requiredCountToolTogenerate=countForCreation.get(key.toString());
                sb.setLength(0);
                for (int i = 1; i <= requiredCountToolTogenerate; i++) {
                    sb.append("CREATE (t")
                      .append(toolGeneratorPlaceHolder)
                      .append(":TOOL {\n")
                      .append("    toolName: 'T_N_").append(domain.substring(0,3)).append(j.substring(0,3)).append('-')
                      .append(i)
                      .append("',\n")
                      .append("    likes: 0,\n")
                      .append("    communicationCount: 0,\n")
                      .append("    deliveryFormat: ''\n")
                      .append("})\n");
                    toolGeneratorPlaceHolder=toolGeneratorPlaceHolder+1;
                }
                System.out.println(sb.toString());
        }
        arrIndex=arrIndex+1;
        }
    }
}

  ----------------------
// Get list of domains and subdomains
  MATCH (N1:SUBDOMAIN)-[:PART_OF]->(N2:DOMAIN)
  WITH N2.domainName AS domainName, COLLECT(N1.subDomainName) AS subDomains
  RETURN {
    domainName: domainName,
    subDomains: subDomains
    } AS result
// Get domain and countOfTools Created under it
    MATCH (N1:TOOL)-[:USED_FOR]->(N2:SUBDOMAIN)
    MATCH (N2:SUBDOMAIN)-[:PART_OF]->(N3:DOMAIN)
    WITH N3.domainName AS domainName, COUNT(N1) AS countOfTools
     ORDER BY countOfTools DESC
    RETURN {
      domainName:domainName,
      countOfTools:countOfTools
    }
// Get the DashBoard
    MATCH (N1:TOOL)-[:USED_FOR]->(N2:SUBDOMAIN)
    MATCH (N2:SUBDOMAIN)-[:PART_OF]->(N3:DOMAIN)
    WITH N3, COUNT (N1) AS countOfTools
     ORDER BY countOfTools DESC
    LIMIT 5
    WITH COLLECT(N3.domainName) AS Top5Domains
    MATCH (N1:TOOL)-[:USED_FOR]->(N2:SUBDOMAIN)
    MATCH (N2:SUBDOMAIN)-[:PART_OF]->(N3:DOMAIN)
    WHERE N3.domainName IN Top5Domains
    WITH N1.deliveryFormat AS deliveryFormat, N3, COUNT(N1.deliveryFormat) AS groupByDeliveryFormat
    RETURN {
      deliveryFormat:deliveryFormat,
      domainName:N3.domainName,
      groupByDeliveryFormat:groupByDeliveryFormat
      } AS RES
// Refresh the Dashboard
      WITH ['DevOps'] AS Top5Domains
      MATCH (N1:TOOL)-[:USED_FOR]->(N2:SUBDOMAIN)
      MATCH (N2:SUBDOMAIN)-[:PART_OF]->(N3:DOMAIN)
      WHERE N3.domainName IN Top5Domains
      WITH N1.deliveryFormat AS deliveryFormat, N3, COUNT(N1.deliveryFormat) AS groupByDeliveryFormat
      RETURN {
        deliveryFormat:deliveryFormat,
        domainName:N3.domainName,
        groupByDeliveryFormat:groupByDeliveryFormat
        } AS RES
        
//-------------------------------
// Tool Page Queries
// Horizontal  list
        MATCH (N1:TOOL)-[:USED_FOR]->(N2:SUBDOMAIN)
        MATCH (N2:SUBDOMAIN)-[:PART_OF]->(N3:DOMAIN)
        WITH N1, N2, N3
        MATCH (N1)
        RETURN {
          toolName:N1.toolName,
          subDomain:N2.subDomainName,
          domain:N3.domainName,
          communicationCount:N1.communicationCount
          } AS res
           ORDER BY N1.communicationCount DESC
          LIMIT 10;

// Tool Dashboard
MATCH (N:TOOL)
RETURN {
  toolName:N.toolName,
  likes:N.likes
  } AS res
ORDER BY N.likes DESC
LIMIT 15;

//  Selected Tools Refresh
            WITH ['T_N_SofBac-10'] AS selectedTools
            MATCH (N:TOOL)
            WHERE N.toolName IN selectedTools
            RETURN {
              toolName:N.toolName,
              likes:N.likes
              } AS res
               ORDER BY N.likes DESC
              LIMIT 15;

//  Menu List
              MATCH (N1:TOOL)-[:USED_FOR]->(N2:SUBDOMAIN)
              MATCH (N2:SUBDOMAIN)-[:PART_OF]->(N3:DOMAIN)
              WITH N3.domainName AS domainName, COLLECT( DISTINCT (N2.subDomainName)) AS subDomain, COLLECT(N1.toolName) AS toolNames
              RETURN {
                domainName:domainName,
                subDomains: subDomain,
                toolsName: toolsName
                } AS result


                // Ongoing
MATCH (N:TOOL{toolName:'T_N_NetArc-1'})
WITH collect(N) as tools
UNWIND range(0,size(tools)-1) as idx
WITH tools[idx] as node
UNWIND range(1,node.communicationCount) as communicationCount
CREATE (c:COMMUNICATION {index: communicationCount, })
RETURN c