# Engineering & Technology Organization Chart

This chart illustrates the structural hierarchy within the Engineering & Technology department, detailing reporting lines, functional verticals, and team composition.

```mermaid
graph TD
    %% Styling Definitions
    classDef executive fill:#E1F5FE,stroke:#01579B,stroke-width:3px,color:#01579B;
    classDef director fill:#E0F2F1,stroke:#00695C,stroke-width:2px,color:#004D40;
    classDef manager fill:#F3E5F5,stroke:#7B1FA2,stroke-width:2px,color:#4A148C;
    classDef individual fill:#FFF3E0,stroke:#E65100,stroke-width:1px,color:#BF360C;
    classDef squad fill:#F9FBE7,stroke:#827717,stroke-width:1px,stroke-dasharray: 5 5,color:#33691E;

    %% Executive Leadership
    VP_ENG[<a href='vp_of_engineering.md'>VP of Engineering</a>]:::executive

    %% Application Development Vertical
    subgraph AppDev_Vertical [Application Development]
        direction TB
        DIR_APP[<a href='director_of_app_dev.md'>Director of App Dev</a>]:::director
        EM_APP[<a href='engineering_manager.md'>Engineering Manager</a>]:::manager

        subgraph App_Team [Product Engineering Team]
            SWE[<a href='software_engineer.md'>Software Engineer</a>]:::individual
            FE[<a href='frontend_engineer.md'>Frontend Engineer</a>]:::individual
            BE[<a href='backend_engineer.md'>Backend Engineer</a>]:::individual
            MOB[<a href='mobile_engineer.md'>Mobile Engineer</a>]:::individual
        end
    end

    %% Infrastructure & Operations Vertical
    subgraph Infra_Vertical [Infrastructure & Operations]
        direction TB
        DIR_INFRA[<a href='director_of_infrastructure.md'>Director of Infrastructure</a>]:::director
        EM_INFRA[<a href='engineering_manager.md'>Engineering Manager</a>]:::manager

        subgraph Infra_Team [Platform Team]
            SRE[<a href='site_reliability_engineer.md'>Site Reliability Engineer</a>]:::individual
            SEC[<a href='security_engineer.md'>Security Engineer</a>]:::individual
            DATA[<a href='data_engineer.md'>Data Engineer</a>]:::individual
            QA[<a href='qa_engineer.md'>QA Engineer</a>]:::individual
        end
    end

    %% Connections
    VP_ENG --> DIR_APP
    VP_ENG --> DIR_INFRA

    DIR_APP --> EM_APP
    EM_APP --> SWE
    EM_APP --> FE
    EM_APP --> BE
    EM_APP --> MOB

    DIR_INFRA --> EM_INFRA
    EM_INFRA --> SRE
    EM_INFRA --> SEC
    EM_INFRA --> DATA
    EM_INFRA --> QA

    %% Cross-Functional Collaboration (Dotted Lines)
    QA -.-> |Quality Assurance| App_Team
    SEC -.-> |Security Audits| App_Team
    DATA -.-> |Data Pipelines| BE
    SRE -.-> |Deployment & Ops| App_Team
```

## Department Structure Overview

The **Engineering & Technology** department is structured to balance product delivery with technical excellence. It is divided into two main verticals:

1.  **Application Development**: Focused on building user-facing products, mobile applications, and the core business logic. This vertical is driven by speed, user experience, and feature delivery.
2.  **Infrastructure & Operations**: Focused on reliability, security, scalability, and data integrity. This vertical enables the App Dev team to move fast by providing stable platforms and tools.

Each vertical is led by a **Director** who reports to the **VP of Engineering**. **Engineering Managers** oversee specific teams or squads within these verticals, ensuring individual contributors (ICs) have the support and resources they need to succeed.
