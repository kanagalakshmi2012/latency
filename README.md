# latency

**OPTIMIZING DNS LOOKUP LATENCY FOR ENHANCED PERFORMANCE IN CLUSTERED SYSTEMS**
* Author: Kanagalakshmi Murugan
* Published In : ConfONE
* Publication Date: 06-2022
* E-ISSN: 2229-7677
* Impact Factor: 9.88
* Link:

**Abstract:**
In distributed clustered systems, efficient network resource resolution is vital for maintaining performance. DNS lookup with a database backend can cause latency and scalability challenges as cluster size grows, due to database query overhead, increased load, and contention. In contrast, host table lookups scale better with lower latency because they use simple in-memory mappings. However, database-backed DNS lookups degrade noticeably with larger clusters, impacting system responsiveness. This paper addresses these issues by optimizing DNS resolution through the use of host tables.

**Key Contributions:**
* Algorithm Development
  Designed and optimized hsot table lookup process to achieve high DNS resolution.
* Performance Comparison
  Conducted bench marking between DataBase DNS lookup time and Hsot Table lookup time.
* Reserach Leadership
  Led the research and technical implementation , focusing on advancing distributed database through algorithm innovation.

**Relevance & Real-World Impact**

* **Efficient DNS Resolution Mechanisms :**\
    DNS resolution was significantly improved, resulting in reduced query latency and faster domain name lookups. These enhancements contributed to more responsive and reliable performance across distributed environments.
* **Enhanced Query Performance:** \
    Overall query latency was significantly reduced as a result of optimized DNS handling.Improved resolution pathways also boosted throughput, enabling faster and more reliable service discovery across distributed systems.need to add here
* **Academic Recognition :** \
    Referenced across academic works and technical analyses related to DNS resolution and query performance enhancements.
* **Educational Impact:** \
    Findings incorporated into educational programs and research efforts, supporting continued academic dialogue on container orchestration and cloud system efficiency.

**Experimental Results (Summary)**:

  | Nodes | Database DNS Lookup Time - Cached (ms) | Host Table Lookup Time (ms) | Improvement (%) |
  |-------|----------------------------------------| ----------------------------| ----------------|
  | 3     | 0.8                                    | 0.1                         | 87.5            |
  | 5     | 0.8                                    | 0.15                        | 81.3            |
  | 7     | 0.9                                    | 0.2                         | 77.8            |
  | 9     | 0.9                                    | 0.25                        | 72.2            |
  | 11    | 0.9                                    | 0.3                         | 66.7            |

**Citation** \
OPTIMIZING DNS LOOKUP LATENCY FOR ENHANCED PERFORMANCE IN CLUSTERED SYSTEMS. \
Kanagalakshmi Murugan \
International Journal on Science and Technology \
E-ISSN-2229-7677 \
**License** \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijsat.org/ \
**Author Contact** \
**LinkedIn**: http://linkedin.com/in/kanagalakshmi-murugan-12b278199 | **Email**: kanagalakshmi2004@gmail.com
